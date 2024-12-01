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

for ($y=0; $y -lt $Map.Count; $y++) {

    $XCoords = $Map[$y] -split '' | Where-Object {$_ -match '\d'}
    for ($x=0; $x -lt $XCoords.Count; $x++) {

        # Found the start of a new basin
        if ($Map[$y][$x] -ne 9 -and $BasinCoords -notcontains "($x,$y)") {

            # The iteration is top down, so we should always be finding the topmost point in the basin
            $AllowedX = @($x)
            $BasinSize = 1
            $BasinCoords += "($x,$y)"
            $BasinsMap[$y][$x] = '.'
            if ($y -eq $Map.Count -1) {$down = $false} else {$down = $true}
            if ($x -eq $XCoords.Count -1 ) {$right -eq $false} else {$right = $true}
            if ($x -eq 0) {$left = $false} else {$left = $true}

            $dx = 1
            $dy = 0
            
            while ($down -or $right -or $left) {

                Format-Basins

                if ( $right -and $BasinsMap[$y+$dy][$x+$dx] -ne '.' -and ($x + $dx) -le ($XCoords.Count - 1) -and ($x + $dx -lt $XCoords.Count) -and (0..8 -contains $Map[$y + $dy][$x + $dx])) { $BasinSize += 1; $AllowedX += ($x + $dx); $BasinCoords += "($($x + $dx),$($y + $dy))"; $BasinsMap[$y+$dy][$x+$dx] = '.' }
                else { $right = $false }

                if ($left -and $BasinsMap[$y+$dy][$x-$dx] -ne '.' -and $dx -ne 0 -and ($x + $dx) -gt 0 -and (0..8 -contains $Map[$y + $dy][$x - $dx])) { $BasinSize += 1; $AllowedX += ($x - $dx); $BasinCoords += "($($x - $dx),$($y + $dy))" ; $BasinsMap[$y+$dy][$x-$dx] = '.' }
                else { $left = $false }

                if (!($right -or $left)) {
                    $dy += 1
                    if (($y + $dy) -lt $Map.Count) {
                        $FoundNewX = $false
                        foreach ($xnew in $AllowedX) {
                            if (0..8 -contains $Map[$y + $dy][$xnew]) {
                                $dx = $xnew - $x
                                $BasinSize += 1
                                $BasinCoords += "($($x + $dx)), $($y + $dy)"
                                $BasinsMap[$y + $dy][$x + $dx] = '.'
                                $FoundNewX = $true
                                break
                            }
                        }
                        if (!($FoundNewX)) { $down = $false }
                        else { $right = $true; $left = $true; $AllowedX = @() }
                    }
                    else { $down = $false }
                }
                else { $dx += 1 }

            }

            $Basins += $BasinSize

            Format-Basins
            Write-Host ""

        }

    }

}

Format-Basins