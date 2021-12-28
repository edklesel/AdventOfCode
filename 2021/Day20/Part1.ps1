$Data = Get-Content "$PSScriptRoot/Input.txt"

$Key = $Data[0]

$Picture = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()
foreach ($Row in $Data[2..($Data.Count)]) { $Picture.Add([System.Collections.Generic.List[String]]($Row -split '' | Where-Object {$_ -in '.','#'})) }

$Steps = 2
$Buffer = 1

# From the key, we need to work out what the behaviour of the pixels outside the picture will be. For the first buffering, the new character will always be .
$Chars = @{'0'='.'}
# As they start as a . if key[0] is . then they will always remain a .
if ($Key[0] -eq '.') {
    $Chars['Odd']='.'
    $Chars['Even']='.'
}
# If key[0] = # then after step 1, they will flip to #
if ($Key[0] -eq '#') {
    $Chars['Odd']='#'
    # Then, if key[511] is . they will flip back
    if ($Key[511] -eq '.') {$Chars['Even'] = '.'}
    # Otherwise, they stay # and the number of pixels is infinite
    else { $Chars['Even'] = '#' }
}
# The area which always flips is the pixels which lie outside a 1 pixel buffer of the starting picture, therefore we only need to care about them

function Format-Picture {

    param([System.Collections.Generic.List[System.Collections.Generic.List[String]]]$Picture)

    $Output = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()

    $local:Buffer = $global:Buffer
    $GridRows = $Picture.Count
    $GridCols = $Picture[0].Count

    if     ($Global:Step -eq 0)     {$Char = $global:Chars['0']}
    elseif ($Global:Step % 2 -eq 1) {$Char = $global:Chars['Odd']}
    elseif ($Global:Step % 2 -eq 0) {$Char = $global:Chars['Even']}

    # Do the top buffer row(s)
    for ($Row = 0; $Row -lt $Buffer; $Row++) {
        $Output.Add([System.Collections.Generic.List[String]]@())
        for ($i=0; $i -lt $GridCols+(2*$Buffer); $i++) { $Output[$Row].Add($Char) }
    }

    for ($Row=0; $Row -lt $GridRows; $Row++) {
        # Add the dot(s) at the start
        $Output.Add([System.Collections.Generic.List[String]]@())
        for ($i=0; $i -lt $Buffer; $i++) {$Output[$Row + $Buffer].Add($Char)}
        # Add the row from the picture
        for ($i=0; $i -lt $GridCols; $i++) {$Output[$Row + $Buffer].Add($Picture[$Row][$i])}
        # Add the dot(s) at the end
        for ($i=0; $i -lt $Buffer; $i++) {$Output[$Row + $Buffer].Add($Char)}
    }

    # Add the final buffer row(s)
    for ($Row = $GridRows+$Buffer; $Row -lt $GridRows + (2*$Buffer); $Row++) {
        $Output.Add([System.Collections.Generic.List[String]]@())
        for ($i=0; $i -lt $GridCols+(2*$Buffer); $i++) { $Output[$Row].Add($Char) }
    }

    return $Output

}

function ConvertFrom-Binary {
    param([String]$BinaryString)
    $Output = 0
    $Split = $BinaryString -split '' | Where-Object {$_ -match '[01]'}
    for ($i=0; $i -lt $Split.Count; $i++) { $Output += [Int32]$Split[$i] * [Math]::Pow(2,($Split.Count - $i -1)) }
    return $Output
}

function Show-Picture { $global:Picture | foreach-object {Write-Output ($_ -join '')} }


# Make the cursor the bottom-right
for ($Step=0; $Step -lt $Steps; $Step++) {

    Write-Host "$($Step+1)/$Steps"

    $Picture = Format-Picture -Picture $Picture

    $NextPicture = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()
    for ($i=0; $i -lt $Picture.Count; $i++) {
        $NextPicture.Add([System.Collections.Generic.List[String]]@())
        for ($j=0; $j -lt $Picture[0].Count; $j++) { $NextPicture[$i].Add('.') }
    }

    for ($x = 0; $x -lt $Picture[0].Count; $x++) {
        for ($y = 0; $y -lt $Picture.Count; $y++) {

            $BinaryString = ''
            for ($yy=($y-1); $yy -le ($y+1); $yy++) {
                for ($xx = ($x-1); $xx -le ($x+1); $xx++) {

                    # If xx or yy are outside the picture, the bit is a .
                    if (0..$($Picture[$y].Count -1) -notcontains $xx -or 0..$($Picture.Count -1) -notcontains $yy  ) {
                        if     ($Global:Step -eq 0)     {$Char = $global:Chars['0']}
                        elseif ($Global:Step % 2 -eq 1) {$Char = $global:Chars['Odd']}
                        elseif ($Global:Step % 2 -eq 2) {$Char = $global:Chars['Even']}
                        $BinaryString = "${BinaryString}$Char"
                    }
                    # Otherwise, the bit comes from the picture
                    else { $BinaryString = "${BinaryString}$($Picture[$yy][$xx])" }
                }
            }

            $BinaryString = $BinaryString -replace '\.','0'
            $BinaryString = $BinaryString -replace '\#','1'

            $Index = ConvertFrom-Binary $BinaryString

            $NextPicture[$y][$x] = $Key[$Index]

        }
    }

    $Picture = $NextPicture

}

$Total = 0
for ($i=0; $i -lt $Picture.Count; $i++) { $Total += ($Picture[$i] | Where-Object {$_ -eq '#'}).Count }
Write-Host "Total Lit Pixels: $Total"