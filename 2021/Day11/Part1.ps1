$Data = Get-Content "$PSScriptRoot/Input.txt"

$Grid = [System.Collections.Generic.List[System.Collections.Generic.List[Int16]]]@()

for ($i=0; $i -lt $Data.Count; $i++) {

    [System.Collections.Generic.List[Int16]]$Row = $Data[$i] -split '' | Where-Object {$_ -match '\d'}
    $Grid.Add($Row)

}

function Format-Grid {
    for ($i=0; $i -lt $global:Grid.Count; $i++) {
        Write-Output "$(($Grid[$i] | ForEach-Object {if ([String]$_.Length -eq 1) {"$_"} else {"$_"}}) -join '')"
    }
    
}

function New-Flash {

    param([Int32]$x,[Int32]$y)

    $global:Flashes ++

    $global:Grid[$y][$x] = 0

    foreach ($Coord in (Get-Coords -x $x -y $y)) {

        if ($Grid[$Coord['y']][$Coord['x']] -ne 0) {$Grid[$Coord['y']][$Coord['x']]++}
        if ($global:Grid[$Coord['y']][$Coord['x']] -gt 9) {New-Flash -x $Coord['x'] -y $Coord['y']}

    }

}

function Get-Coords {

    param([Int32]$x,[Int32]$y)

    $Coords = @()
    $Coords += @{'x'=$x;'y'=$y}

    # Horizontal
    if ($x -ne $global:Grid[0].Count -1) {$Coords += @{'x'=$x+1;'y'=$y}}
    if ($x -ne 0) {$Coords += @{'x'=$x-1;'y'=$y}}

    # Vertical
    if ($y -ne $global:Grid.Count -1) {$Coords += @{'x'=$x;'y'=$y+1}}
    if ($y -ne 0) {$Coords += @{'x'=$x;'y'=$y-1}}

    # Top left
    if ($y -ne 0 -and $x -ne 0) {$Coords += @{'x'=$x-1; 'y'=$y-1}}
    # Top right
    if ($y -ne 0 -and $x -ne $global:Grid[0].Count -1) {$Coords += @{'x'=$x+1; 'y'=$y-1}}
    # Bottom right
    if ($y -ne $global:Grid.Count -1 -and $x -ne $global:Grid[0].Count -1) {$Coords += @{'x'=$x+1; 'y'=$y+1}}
    # Bottom left
    if ($y -ne $global:Grid.Count -1 -and $x -ne 0) {$Coords += @{'x'=$x-1; 'y'=$y+1}}

    return $Coords
}

$Flashes = 0
$Steps = 100
for ($Step=1; $Step -le $Steps; $Step++) {
    
    # Increase the energy by 1
    for ($y=0; $y -lt $Grid.Count; $y++) {
        for ($x=0; $x -lt $Grid[$y].Count; $x++) {
            $Grid[$y][$x]++
        }
    }

    # Flash any octopuses
    for ($y=0; $y -lt $Grid.Count; $y++) {
        for ($x=0; $x -lt $Grid[$y].Count; $x++) {
            if ($Grid[$y][$x] -gt 9) {New-Flash -x $x -y $y}
        }
    }

}

Write-Host "After step ${Step}: $Flashes flashes."
