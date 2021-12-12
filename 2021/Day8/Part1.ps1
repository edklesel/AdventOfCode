[System.Collections.ArrayList]$Data = Get-Content "$PSScriptRoot/Input.txt"

$Count = 0

foreach ($Row in $Data) {

    $Split = $Row -split ' \| '

    $Output = $Split[1]

    $Displays = $Output -split ' '

    $Count += ($Displays | where-object {2,3,4,7 -contains ($_ -split '' | Where-Object {$_-match '\w'}).Count}).Count

}

Write-Host "1,4,7,8 appear $Count times in the output displays."