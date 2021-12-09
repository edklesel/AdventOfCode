[System.Collections.ArrayList]$Data = (Get-Content "$PSScriptRoot/Input.txt") -split ','

$Max = $Data | Measure-Object -Maximum | Select-Object -ExpandProperty Maximum
$Min = $Data | Measure-Object -Minimum | Select-Object -ExpandProperty Minimum

$MinFuel = $Max * $Data.Count * $($Data | Measure-Object -Sum | Select-Object -ExpandProperty Sum)

foreach ($Index in $Min..$Max) {
    
    $Index

    $Differences = $Data | ForEach-Object {($([math]::Abs($_ - $Index))*($([math]::Abs($_ - $Index))+1))/2}
    $FuelUsed = $Differences | Measure-Object -Sum | Select-Object -ExpandProperty Sum

    if ($FuelUsed -lt $MinFuel) {$MinFuel = $FuelUsed; $Position = $Index}

}

Write-Host "Minimum Fuel: $MinFuel at position $Position"