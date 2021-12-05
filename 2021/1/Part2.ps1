Set-Location $PSScriptRoot
$Report = Get-Content ./Input.txt

$Index = 1 # Skip the first one
$Larger = 0
while ($Index -le $Report.Count - 2) {

    $Window2 = [Int32]$Report[$Index+2] + [Int32]$Report[$Index+1] + [Int32]$Report[$Index]
    $Window1 = [Int32]$Report[$Index+1] + [Int32]$Report[$Index] + [Int32]$Report[$Index-1]

    if ($Window2 -gt $Window1) {
        $Larger += 1
        Write-Host -NoNewline '+'
    }
    elseif ($Window2 -lt $Window1) {
        Write-Host -NoNewline '-'
    }
    else {Write-Host -NoNewline '='}
    # else {Write-Host "Decreased"}

    $Index += 1

}

Write-Host
Write-Host $Larger