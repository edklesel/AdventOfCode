$Horizontal = 0
$Depth = 0

$Commands = Get-Content "$PSScriptRoot/Input.txt"

foreach ($Command in $Commands) {

    $Direction = ($Command -split ' ')[0]
    $Magnitude = ($Command -split ' ')[1]

    switch ($Direction) {
        'forward' {$Horizontal += $Magnitude}
        'up' {$Depth -= $Magnitude}
        'down' {$Depth += $Magnitude}
    }

}

Write-Host "Horizontal: $Horizontal"
Write-Host "Depth: $Depth"
Write-Host "Product: $($Horizontal * $Depth)"