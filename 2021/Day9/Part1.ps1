$Map = Get-Content -Path "$PSScriptRoot/Input.txt"
for ($y=0; $y -lt $Map.Count; $y++) { $Map[$y] = ($Map[$y] -split '' | Where-Object {$_ -match '\w'}) }

[System.Collections.ArrayList]$LowPoints = @()

for ($y=0; $y -lt $Map.Count; $y++) {

    $up=$true; $down=$true
    if ($y -eq 0) {$up = $false}
    elseif ($y -eq ($Map.Count -1)) {$down = $false}

    $XCoords = $Map[$y] -split '' | Where-Object {$_ -match '\d'}
    for ($x=0; $x -lt $XCoords.Count; $x++) {

        $left=$true; $right=$true
        if ($x -eq 0) {$left = $false}
        elseif ($x -eq ($XCoords.Count -1)) {$right = $false}

        $Vals = @($Map[$y][$x])
        if ($up) { $Vals += $Map[$y-1][$x] }
        if ($down) { $Vals += $Map[$y+1][$x] }
        if ($left) { $Vals += $Map[$y][$x-1] }
        if ($right) { $Vals += $Map[$y][$x+1] }

        if (($Vals | Measure-Object -Minimum).Minimum -eq $Map[$y][$x] -and ($Vals | Where-Object {$_ -eq $Map[$y][$x]}).Count -eq 1) {
            
            $LowPoints.Add([Int]($Map[$y][$x])) | Out-Null

        }

    }

}

Write-Host "Total Risk: $(($LowPoints | Measure-Object -Sum).Sum + $LowPoints.Count)"