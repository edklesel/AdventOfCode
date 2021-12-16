$Data = Get-Content "$PSScriptRoot/Input.txt"
$Template = $Data | Where-Object { $_ -match '\w{3,}' }
$Insertions = $Data | Where-Object { $_ -match '\w{2} \-\> \w' }
$Pairs = @{}

# Work out the number of pairs
for ($i=0; $i -lt ($Template.Length -1); $i++) {
    $Pair = "$($Template[$i])$($Template[$i+1])"
    if ($Pairs.Keys -notcontains $Pair) {$Pairs[$Pair] = 0}
    $Pairs[$Pair]++
}
# break
$Step = 0
$Steps = 40

while ($Step -lt $Steps) {

    $NewPairs = @{}

    foreach ($Insertion in $Insertions) {

        $Search = ($Insertion -split ' \-\> ')[0]
        $Insert = ($Insertion -split ' \-\> ')[1]
        $NewPairL = "$($Search[0])$($Insert.ToLower())"
        $NewPairR = "$($Insert.ToLower())$($Search[1])"

        if (($Pairs.Keys -notcontains $Search) -or ($Pairs[$Search] -eq 0)) {continue}
        
        if ($NewPairs.Keys -notcontains $NewPairL) {$NewPairs[$NewPairL] = 0}
        if ($NewPairs.Keys -notcontains $NewPairR) {$NewPairs[$NewPairR] = 0}
        $NewPairs[$NewPairL] += $Pairs[$Search]
        $NewPairs[$NewPairR] += $Pairs[$Search]
        $Pairs[$Search] = 0

    }

    foreach ($Pair in $NewPairs.Keys) {

        $NewPair = $Pair.ToUpper()
        if ($Pairs.Keys -notcontains $NewPair) { $Pairs[$NewPair] = 0 }
        $Pairs[$NewPair] = $NewPairs[$NewPair]

    }

    $Step++

}

$Occurrences = @{}
foreach ($Pair in $Pairs.Keys) {
    foreach ($Char in ($Pair -split '' | Where-Object {$_ -match '\w'})) {
        if ($Occurrences.Keys -notcontains $Char) {$Occurrences[$Char] = 0}
        $Occurrences[$Char] += $Pairs[$Pair]
    }
}

$NewOccurrences = @{}
foreach ($Char in $Occurrences.Keys) {
    if ($Occurrences[$Char] % 2 -eq 0) {$NewOccurrences[$Char] = $Occurrences[$Char]/2}
    else {$NewOccurrences[$Char] = ($Occurrences[$Char]+1)/2}
}

$Min = [Math]::Pow(10,10000)
$Max = 0
foreach ($Char in $NewOccurrences.Keys) {
    if ($NewOccurrences[$Char] -gt $Max) {$Max = $NewOccurrences[$Char]}
    elseif ($NewOccurrences[$Char] -lt $Min) {$Min = $NewOccurrences[$Char]}
}

Write-Host "Difference after $Steps steps: $($Max - $Min)"