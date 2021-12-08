$Data = Get-Content "$PSScriptRoot/Input.txt"

# Create the grid
$GridSize = 1000
$Grid = [System.Collections.ArrayList]@()
for ($i=0; $i -lt $GridSize; $i++) {
    $Zeroes = '0'*$GridSize -split '' | ForEach-Object {[Int32]$_}
    $Grid.Add($Zeroes) | Out-Null
}

function Format-Grid {
    param ($Grid)

    $Output = ''

    for ($i=0; $i -lt $global:GridSize; $i++) {
        for ($j=0; $j -lt $global:GridSize; $j++) {
            $Output = "${Output}$($global:Grid[$j][$i] -replace '0','.')"
        }
        $Output = "${Output}`n"
    }

    return $Output

}

function New-Vent {
    param(
        [Int32]$x1,
        [Int32]$y1,
        [Int32]$x2,
        [Int32]$y2
    )

    # Diagonal
    if ($x1 -ne $x2 -and $y1 -ne $y2) {
        $x_range = $x1..$x2
        $y_range = $y1..$y2
    }
    # Vertical
    elseif ($x1 -eq $x2) {
        $y_range = $y1..$y2
        $x_range = $y_range | Foreach-Object {$x1}
    }
    # Horizontal
    elseif ($y1 -eq $y2) {
        $x_range = $x1..$x2
        $y_range = $x_range | ForEach-Object {$y1}
    }

    $i = 0
    while ($i -lt $x_range.Count) {
        $x = $x_range[$i]
        $y = $y_range[$i]
        $global:Grid[$x][$y]++
        $i++
    }
    
}

foreach ($Row in $Data) {

    $Raw1 = $Row -split ' -> '

    $x1 = ($Raw1[0] -split ',')[0]
    $y1 = ($Raw1[0] -split ',')[1]
    $x2 = ($Raw1[1] -split ',')[0]
    $y2 = ($Raw1[1] -split ',')[1]

    New-Vent -x1 $x1 -y1 $y1 -x2 $x2 -y2 $y2
    
}

Write-Host "Intersections: $(($Grid | foreach-Object {$_ | Where-Object {$_ -gt 1}}).Count)"