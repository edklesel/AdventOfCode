$Data = Get-Content -Path "$PSScriptRoot/Input.txt"

$Hex = @{
    '0' = '0000'
    '1' = '0001'
    '2' = '0010'
    '3' = '0011'
    '4' = '0100'
    '5' = '0101'
    '6' = '0110'
    '7' = '0111'
    '8' = '1000'
    '9' = '1001'
    'A' = '1010'
    'B' = '1011'
    'C' = '1100'
    'D' = '1101'
    'E' = '1110'
    'F' = '1111'
}

function ConvertFrom-Binary {
    param($BinaryString)

    [Array]$BinaryString = $BinaryString -split '' | Where-Object {$_ -match '\d'}

    $Total = 0
    for ($i=0; $i -lt $BinaryString.Count; $i++) {

        $Power = $BinaryString.Count - $i -1
        $Total += [Int64]($BinaryString[$i]) * [Math]::Pow(2,$Power)

    }

    return $Total
}

function Get-Packet {
    param(
        [System.Collections.Generic.List[Int16]]$Packet
    )
    $Header_Version = ConvertFrom-Binary $Packet[0..2]

    $global:VersionTotal += $Header_Version

    $Header_Type = ConvertFrom-Binary $Packet[3..5]

    $Body = $Packet[6..($Packet.Count-1)]

    # If the packet is a literal packet
    if ($Header_Type -eq 4) {

        $Output = ''
        $i = 0
        while ($true) {

            $String = $Body[$i..($i+4)] -join ''
            $Output = "${Output}$($String[1..4] -join '')"

            if ($String[0] -eq '0') {break}

            $i += 5

        }

        if ($i+5 -eq $Body.Count) { $RemainingBody = '' }
        else { $RemainingBody = $Body[($i+5)..$($Body.Count -1)] }

    }

    # Otherwise, if the packet is an operator packet
    else {

        $LengthTypeID = $Body[0]

        if ($LengthTypeID -eq 0) {

            $PacketLength = ConvertFrom-Binary -BinaryString $Body[1..15]

            $SubPackets = $Body[16..$($PacketLength + 15)]

            while ($SubPackets -contains '1') {

                $SubPackets = Get-Packet -Packet $SubPackets

            }

            if (($PacketLength + 16) -eq $Body.Count) { $RemainingBody = $null }
            else { $RemainingBody = $Body[$($PacketLength + 16)..($Body.Count -1)] }

        }
        elseif ($LengthTypeID -eq 1) {

            $PacketCount = ConvertFrom-Binary -BinaryString $Body[1..11]

            $i = 0
            $SubPackets = $Body[12..$($Body.Count -1)]

            while ($i -lt $PacketCount) {
                $SubPackets = Get-Packet -Packet $SubPackets
                $i++
            }

            $RemainingBody = $SubPackets

        }

    }

    if ($RemainingBody -contains '1') { $ReturnValue = $RemainingBody }
    else { $ReturnValue = $null }

    return $ReturnValue

}

# Construct the binary character
$Binary = [System.Collections.Generic.List[Int16]]@()
foreach ($Char in ($Data -split '' | Where-Object {$_ -match '\w'})) {
    $Hex[$Char] -split '' | Where-Object { $_ -match '\d' } | foreach-object {$Binary.Add($_)}
}

$VersionTotal = 0

Get-Packet -Packet $Binary

Write-Host "Sum of all packet versions: $VersionTotal"