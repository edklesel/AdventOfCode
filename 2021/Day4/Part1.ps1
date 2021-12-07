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
            foreach ($RowNum in 1..5) { $Value = "${Value}$($global:Boards[$BoardNum][$RowNum])" }
            if ($Value -eq 'xxxxx') { $Winner = $BoardNum; break }
        }

        # Check Diagonal TL->BR
        $Value = ''
        for ($i=0; $i -lt $global:Boards[$BoardNum][0].Count; $i++) { $Value = "${Value}$($global:Boards[$BoardNum][$i][$i])" }
        if ($Value -eq 'xxxxx') { $Winner = $BoardNum; break }

        # Check Diagonal TR->BL
        $Value = ''
        for ($i=0; $i -lt $global:Boards[$BoardNum][0].Count; $i++) { $Value = "${Value}$($global:Boards[$BoardNum][$i][$($global:Boards[$BoardNum][0].Count)-$i])" }
        if ($Value -eq 'xxxxx') { $Winner = $BoardNum; break }

    }

    return $Winner
}

foreach ($Number in $RandomNumbers) {

    Test-Number -Number $Number

    $Winner = Test-Board

    if ($null -ne $Winner) { break }

}

$WinningBoard = $Boards[$Winner]

$Sum = 0
foreach ($Row in $WinningBoard) {
    foreach ($Value in $Row) {
        if ($Value -ne 'x') {$Sum += $Value}
    }
}

Write-Host "Score: $([Int32]$Number * [Int32]$Sum)"