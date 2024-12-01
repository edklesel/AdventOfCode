$Data = Get-Content "$PSScriptRoot/Input.txt"

$Cubes = [System.Collections.Generic.List[String]]@()

$Region = @{'xmin'=-50;'xmax'=50;'ymin'=-50;'ymax'=50;'zmin'=-50;'zmax'=50;}

foreach ($Row in $Data) {

    $null = $Row -match '(?<Switch>\w+) x=(?<xmin>\-?\d+)\.\.(?<xmax>\-?\d+)\,y=(?<ymin>\-?\d+)\.\.(?<ymax>\-?\d+)\,z=(?<zmin>\-?\d+)\.\.(?<zmax>\-?\d+)'
    $Switch = $matches['switch']
    $xmin = [Int32]$matches['xmin']
    $xmax = [Int32]$matches['xmax']
    $ymin = [Int32]$matches['ymin']
    $ymax = [Int32]$matches['ymax']
    $zmin = [Int32]$matches['zmin']
    $zmax = [Int32]$matches['zmax']

    foreach ($x in ($xmin..$xmax | Where-Object {$_ -ge $Region['xmin'] -and $_ -le $Region['xmax']})) {
        foreach ($y in $ymin..$ymax | Where-Object {$_ -ge $Region['ymin'] -and $_ -le $Region['ymax']}) {
            foreach ($z in $zmin..$zmax | Where-Object {$_ -ge $Region['zmin'] -and $_ -le $Region['zmax']}) {

                $Cube = "($x,$y,$z)"

                if ($Cubes -match $Cube -and $Switch -eq 'off') { $Cubes.Remove($Cube) | Out-Null }
                elseif (-not ($Cubes -match $Cube) -and $Switch -eq 'on') { $Cubes.Add($Cube) }

            }
        }
    }

    $Cubes.Count
    Write-Host "Row"

}