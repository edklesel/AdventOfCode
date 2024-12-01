$Map = Get-Content -Path "$PSScriptRoot/Input.txt"
for ($y=0; $y -lt $Map.Count; $y++) { $Map[$y] = ($Map[$y] -split '' | Where-Object {$_ -match '\w'}) }

$BasinsMap = Get-Content -Path "$PSScriptRoot/Input.txt"
for ($y=0; $y -lt $BasinsMap.Count; $y++) { $BasinsMap[$y] = ($BasinsMap[$y] -split '' | Where-Object {$_ -match '\w'}) }

$BasinCoords = [System.Collections.ArrayList]@()
$Basins = [System.Collections.ArrayList] @()

function Format-Basins {    
    Remove-Item -ErrorAction SilentlyContinue "$PSScriptRoot/Basins.txt"
    for ($y=0; $y -lt $global:BasinsMap.Count; $y++) { $global:BasinsMap[$y] -join '' }
}

function Check-BelowRow {

    param(
        [Int32]$y,[Int32]$x
    )

    $dx = 0
    while ($true) {
        
    }

        
    if ($global:Maps[$y+1][$x + $dx] -ne 9) {



    }

}

for ($y=0; $y -lt $Map.Count; $y++) {

    $XCoords = $Map[$y] -split '' | Where-Object {$_ -match '\d'}
    for ($x=0; $x -lt $XCoords.Count; $x++) {

        # Found the start of a new basin
        if ($Map[$y][$x] -ne 9 -and $BasinCoords -notcontains "($x,$y)") {


            
        }

    }

}

Format-Basins