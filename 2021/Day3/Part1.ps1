$Report = Get-Content "$PSScriptRoot/Input.txt"

$NumBits = $Report[0] | Measure-Object -Character | Select-Object -ExpandProperty Characters

$Bits = @()
foreach ($i in 1..$NumBits) {
    $Bits += @{'0'=0;'1'=0}
}

foreach ($Item in $Report) {

    foreach ($Index in 0..$($NumBits-1)) {
        $Bits[$Index]["$(($Item -split '')[$Index+1])"] += 1
    }

}

$GammaRate = ''
$EpsilonRate = ''

foreach ($Bit in $Bits) {

    switch ($Bit['0'] -lt $Bit['1']) {
        $True  {$GammaRate = "${GammaRate}1"; $EpsilonRate = "${EpsilonRate}0"}
        $False {$GammaRate = "${GammaRate}0"; $EpsilonRate = "${EpsilonRate}1"}
    }
    
}

$GammaRate_Int = 0
$EpsilonRate_Int = 0

foreach ($Index in 0..$($NumBits-1)) {

    $GammaRate_Int += [Int]($GammaRate -split '')[$NumBits-$Index] * [math]::Pow(2,$Index)
    $EpsilonRate_Int += [Int]($EpsilonRate -split '')[$NumBits-$Index] * [math]::Pow(2,$Index)

}

Write-Host "Power Consumption: $($GammaRate_Int * $EpsilonRate_Int)"