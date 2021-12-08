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

    if ($x1 -ne $x2 -and $y1 -ne $y2) {

        $x_range = $x1..$x2
        $y_range = $y1..$y2
        $i = 0
        while ($i -lt $x_range.Count) {
            $x = $x_range[$i]
            $y = $y_range[$i]
            $global:Grid[$x][$y]++
            $i++
        }

    }

    if ($x1 -eq $x2) {

        $x = $x1
        if ($y1 -lt $y2) {$top = $y1; $bottom = $y2}
        else {$top = $y2; $bottom = $y1}

        for ($y=$top; $y -le $bottom; $y++) {
            $global:Grid[$x][$y]++
        }

    }
    elseif ($y1 -eq $y2) {

        $y = $y1
        if ($x1 -lt $x2) {$left = $x1; $right = $x2}
        else {$left = $x2; $right = $x1}

        for ($x=$left; $x -le $right; $x++) {
            $global:Grid[$x][$y]++
        }

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