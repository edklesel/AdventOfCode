$Report = Get-Content "$PSScriptRoot/Input.txt"

$NumBits = $Report[0] | Measure-Object -Character | Select-Object -ExpandProperty Characters

function ConvertFrom-Binary {
    param ([String]$Binary)

    $Binary_arr = $Binary -split '' | Where-Object {$_ -match '\d'}
    $Total = 0

    foreach ($Index in 0..$($Binary_arr.Count - 1)) {$Total += [Int]$Binary_arr[$Binary_arr.Count-$Index-1] * [math]::Pow(2,$Index)}

    return $Total

}

$Bits = @()

function Get-Bits {
    param(
        [System.Collections.ArrayList]$Data
    )

    $NumBits = $Report[0] | Measure-Object -Character | Select-Object -ExpandProperty Characters

    $Bits = @()

    foreach ($i in 1..$NumBits) {
        $Bits += @{'0'=0;'1'=0}
    }
    
    foreach ($Item in $Data) {

        $Item_arr = $Item -split ',' | Where-Object {$_ -match '\d'}
    
        foreach ($Index in 0..$($NumBits-1)) {
            $Bits[$Index]["$($Item_arr[$Index])"] += 1
        }
    
    }

    return $Bits

}

foreach ($i in 1..$NumBits) {
    $Bits += @{'0'=0;'1'=0}
}

foreach ($Item in $Report) {

    $Item_arr = $Item -split ',' | Where-Object {$_ -match '\d'}

    foreach ($Index in 0..$($NumBits-1)) {
        $Bits[$Index]["$($Item_arr[$Index])"] += 1
    }

}

$GammaRate = ''
$EpsilonRate = ''
$Oxy_Bits = $Report.Clone()
$CO2_Bits = $Report.Clone()

$i = 0
foreach ($Bit in $Bits) {

    switch ($Bit['0'] -lt $Bit['1']) {
        $True  {
            $GammaRate = "${GammaRate}1"; $EpsilonRate = "${EpsilonRate}0"
        }
        $False {
            $GammaRate = "${GammaRate}0"; $EpsilonRate = "${EpsilonRate}1"
        }
    }

    $i += 1
    
}

$Index = 0
while ($true) {

    if ($Oxy_Bits.count -eq 1) {break}

    $BitsO = Get-Bits -Data $Oxy_Bits

    if ($BitsO[$Index]["0"] -le $BitsO[$Index]["1"]) {$Oxy_Bits = $Oxy_Bits | Where-Object {$_[$Index] -eq "1"}}
    else {$Oxy_Bits = $Oxy_Bits | Where-Object {$_[$Index] -eq "0"}}

    $Index += 1
}

$Index = 0
while ($true) {

    if ($CO2_Bits.count -eq 1) {break}

    $BitsC = Get-Bits -Data $CO2_Bits

    if ($BitsC[$Index]["0"] -le $BitsC[$Index]["1"]) {$CO2_Bits = $CO2_Bits | Where-Object {$_[$Index] -eq "0"}}
    else {$CO2_Bits = $CO2_Bits | Where-Object {$_[$Index] -eq "1"}}

    $Index += 1
}

$GammaRate_Int = ConvertFrom-Binary -Binary $GammaRate
$EpsilonRate_Int = ConvertFrom-Binary -Binary $EpsilonRate
$OxygenConsumptionRate_Int = ConvertFrom-Binary -Binary $Oxy_Bits
$CO2ScrubberRate_Int = ConvertFrom-Binary -Binary $CO2_Bits

Write-Host "Power Consumption: $($GammaRate_Int * $EpsilonRate_Int)"
Write-Host "Life Support Rating: $($OxygenConsumptionRate_Int * $CO2ScrubberRate_Int)"