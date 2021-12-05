Set-Location $PSScriptRoot
$Report = Get-Content ./Input.txt

$Index = 1 # Skip the first one
$Larger = 0
while ($Index -le $Report.Count) {

    if ([Int32]$Report[$Index] -gt [Int32]$Report[$Index-1]) {
        $Larger += 1
    }
    # else {Write-Host "Decreased"}

    $Index += 1

}

Write-Host $Larger