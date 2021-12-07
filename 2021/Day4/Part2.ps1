$Data = Get-Content "$PSScriptRoot/Input.txt"
$RandomNumbers = $Data[0] -split ','

# Create an array of all the boards
[System.Collections.ArrayList]$Boards = @()
for ($i=2;$i -le $($Data.Count);$i+=6) {$Boards.Add([System.Collections.ArrayList]@($Data[$i..$($i+4)])) | Out-Null}

# Convert the rows into arrays
for ($BoardNum=0; $BoardNum -lt $Boards.Count; $BoardNum++) {
    for ($RowNum=0; $RowNum -lt $Boards[$BoardNum].Count; $RowNum ++) {
        $Boards[$BoardNum][$RowNum] = ($Boards[$BoardNum][$RowNum] -replace '^\s','') -split '\s+'
    }
}

function Test-Number {
    param([String]$Number)
    for ($BoardNum=0; $BoardNum -lt $global:Boards.Count; $BoardNum++) {
        for ($RowNum=0; $RowNum -lt $global:Boards[$BoardNum].Count; $RowNum ++) {
            for ($ColNum=0; $ColNum -lt $global:Boards[$BoardNum][$RowNum].Count; $ColNum++) {
                if ($global:Boards[$BoardNum][$RowNum][$ColNum] -eq $Number) {$global:Boards[$BoardNum][$RowNum][$ColNum] = 'x'}
            }
        }
    }
}

function Test-Board {

    $Winner = $null

    for ($BoardNum=0; $BoardNum -lt $global:Boards.Count; $BoardNum++) {

        # Check Horizontals
        for ($RowNum=0; $RowNum -lt $global:Boards[$BoardNum].Count; $RowNum++) {
            if ($global:Boards[$BoardNum][$RowNum] -join '' -eq 'xxxxx') { $Winner = $BoardNum; break }
        }

        # Check Verticals
        for ($ColNum=0; $ColNum -lt $global:Boards[0].Count; $ColNum++) {
            $Value = ''
            foreach ($RowNum in 0..4) { $Value = "${Value}$($global:Boards[$BoardNum][$RowNum][$ColNum])" }
            if ($Value -eq 'xxxxx') { $Winner = $BoardNum; break }
        }

    }

    return $Winner
}

function Format-Board {
    param($Board)

    $Output = ''

    foreach ($Row in $Board) {
        foreach ($Value in $Row) {

            if ([String]$Value.Length -eq 1) {$Value = " $Value"}

            $Output += "$Value "

        }

        $Output += "`n"
    }

    return $Output
}

foreach ($Number in $RandomNumbers) {

    Test-Number -Number $Number

    # While loop to check if multiple boards are winners when a number is crossed off
    while ($true) {

        $Result = Test-Board

        if ($null -eq $Result) {break}

        if ($null -ne $Result) {

            $Result = Test-Board

            $WinningBoard = $global:Boards[$Result]
            $Sum = 0
            foreach ($Row in $WinningBoard) {
                foreach ($Value in $Row) {
                    if ($Value -ne 'x') {$Sum += $Value}
                }
            }
            $WinningNumber = $Number
            Write-Host "Number: $Number, Score: $([Int32]$WinningNumber * [Int32]$Sum)"
            $global:Boards.RemoveAt($Result)
        }

    }

}
