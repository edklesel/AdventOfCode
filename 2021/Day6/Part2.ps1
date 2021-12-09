$Fishes = @(0,0,0,0,0,0,0,0,0)

(Get-Content "$PSScriptRoot/Input.txt") -split ',' | foreach-object { $Fishes[$_]++ }

$Day = 0
$MaxDays = 256
$Fishes.Keys | Sort-Object {[Int32]$_} | Foreach-Object {$Fishes["$_"]}

while ($Day -lt $MaxDays) {

    $Day ++
    $ZeroFishes = $Fishes[0]

    foreach ($i in 1..9) {
        $Fishes[$i-1] = $Fishes[$i]
    }
    $Fishes[6] += $ZeroFishes
    $Fishes[8] += $ZeroFishes
    # $Fishes | Measure-Object -Sum | Select-Object -ExpandProperty Sum
    # $Day
    # Write-Host ""
    # $Fishes
    # Write-Host ""

}

# $($Fishes | Foreach-Object { $_; $Total += $Fishes[$_] })
Write-Host "Total after Day ${MaxDays}: $($Fishes | Measure-Object -Sum | Select-Object -ExpandProperty Sum)"