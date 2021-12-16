$Data = Get-Content "$PSScriptRoot/Input.txt"
$Coords = $Data | Where-Object { $_ -match '\d+\,\d+' }
$Folds = $Data | Where-Object { $_ -match 'fold along [xy]\=\d+' }

# Work out the size of the grid
$GridX = 0
$GridY = 0
foreach ($Coord in $Coords) {
    if ([Int32]($Coord -split ',')[0] -gt $GridX) {$GridX = [Int32]($Coord -split ',')[0]}
    if ([Int32]($Coord -split ',')[1] -gt $GridY) {$GridY = [Int32]($Coord -split ',')[1]}
}

# Create the grid
$Grid = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()
for ($y=0; $y -le $GridY; $y++) {
    $Grid.Add([System.Collections.Generic.List[String]]@())
    for ($x=0; $x -le $GridX; $x++) {
        $Grid[$y].Add('.')
    }
}

function Format-Grid { foreach ($Row in $global:Grid) { Write-Output $($Row -join '') } }

foreach ($Coord in $Coords) {

    $x = [Int32]($Coord -split ',')[0]
    $y = [Int32]($Coord -split ',')[1]
    $Grid[$y][$x] = '#'
    
}

$FoldCount = 0
foreach ($Fold in $Folds) {

    $FoldCount ++
    $Fold -match 'fold along (?<Axis>\w)=(?<Coord>\d+)' | Out-Null
    $FoldAxis = $matches['Axis']
    $FoldCoord = [Int32]$matches['Coord']

    Write-Host "$FoldAxis fold!"

    switch ($FoldAxis) {
        
        # Fold along the y axis
        'y' {

            for ($y = $FoldCoord + 1; $y -lt $Grid.Count; $y++) {
            
                for ($x = 0; $x -lt $Grid[$y].Count; $x++) {
            
                    # Skip dots
                    if ($Grid[$y][$x] -eq '.') { continue }
            
                    $Grid[(2*$FoldCoord - $y)][$x] = '#'
            
                }
            
            }
            
            $Grid = $Grid[0..($FoldCoord)]
        }

        # Fold along the x axis
        'x' {

            for ($y = 0; $y -lt $Grid.Count; $y++) {

                for ($x = $FoldCoord +1; $x -lt $Grid[$y].Count; $x++) {

                    if ($Grid[$y][$x] -eq '.') { continue }

                    $Grid[$y][2*$FoldCoord - $x] = '#'

                }

            }

            for ($y=0; $y -lt $Grid.Count; $y++) {
                $Grid[$y] = $Grid[$y][0..($FoldCoord - 1)]
            }

        }
    
    }

    $Dots = 0
    foreach ($Row in $Grid) { $Dots += ($Row | Where-Object { $_ -eq '#' }).Count }
    
    Write-Host "After $FoldCount folds: $Dots dots."
    Write-Host ""

}

Format-Grid