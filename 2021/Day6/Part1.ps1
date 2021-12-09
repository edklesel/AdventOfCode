class LanternFish {
    [Int32]$TimeLeft
    [Boolean]$NewFish
}

$Fishes = [System.Collections.ArrayList]@()

function New-LanternFish {
    param (
        [Switch]$NewFish,
        [Int32]$TimeLeft=9
    )
    $LanternFish = New-Object LanternFish
    $LanternFish.TimeLeft = $TimeLeft
    if ($NewFish) {$LanternFish.NewFish = $true}

    $global:Fishes.Add($LanternFish) | Out-Null
}

(Get-Content "$PSScriptRoot/Input.txt") -split ',' | ForEach-Object {
    New-LanternFish -TimeLeft $_
}

$Day = 1
$MaxDays = 80
# Write-Host "Initial State: $($Fishes.TimeLeft -join ',')"
while ($Day -le $MaxDays) {

    for ($i=0; $i -lt $Fishes.Count; $i++) {
        $Fishes[$i].TimeLeft -= 1
        if ($Fishes[$i].TimeLeft -lt 0) {
            New-LanternFish -NewFish
            $Fishes[$i].TimeLeft = 6
            $Fishes[$i].NewFish = $false
        }
    }

    # Write-Host "After Day ${Day}: $($Fishes.TimeLeft -join ',')"

    Write-Host $Day - $($Fishes.Count)
    $Day++
}

Write-Host ""
Write-Host "Fishes after day ${MaxDays}: $($Fishes.Count)"