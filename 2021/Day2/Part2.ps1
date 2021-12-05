$Horizontal = 0
$Depth = 0
$Aim = 0

$Commands = Get-Content "$PSScriptRoot/Input.txt"

foreach ($Command in $Commands) {

    $Direction = ($Command -split ' ')[0]
    $Magnitude = ($Command -split ' ')[1]

    switch ($Direction) {
        'forward' {$Horizontal += $Magnitude; $Depth += $Aim * $Magnitude}
        'up' {$Aim -= $Magnitude}
        'down' {$Aim += $Magnitude}
    }

}

Write-Host "Horizontal: $Horizontal"
Write-Host "Depth: $Depth"
Write-Host "Product: $($Horizontal * $Depth)"