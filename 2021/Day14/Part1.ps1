$Data = Get-Content "$PSScriptRoot/Input.txt"
$Template = $Data | Where-Object { $_ -match '\w{3,}' }
$Insertions = $Data | Where-Object { $_ -match '\w{2} \-\> \w' }

$Step = 0
$Steps = 10

while ($Step -lt $Steps) {

    foreach ($Insertion in $Insertions) {

        $Search = ($Insertion -split ' \-\> ')[0]

        # Put the insert between the two matches, but make it lower so it doesn't match again this step
        $Search_Arr = $Search -split '' | Where-Object { $_ -match '\w' }
        $Insert = "$($Search_Arr[0])$(($Insertion -split ' \-\> ')[1].ToLower())$($Search_Arr[1])"

        $Template = $Template -creplace $Search,$Insert
        while (($Template -creplace $Search,$Insert) -ne $Template) {
            $Template = $Template -creplace $Search,$Insert
        }

    }

    $Template = $Template.ToUpper()
    $Step++

}

$Occurrences = @{}
foreach ($Char in ($Template -split '' | Where-Object {$_ -match '\w'} | Sort-Object | Get-Unique)) {
    $Occurrences[$Char] = ($Template -split '' | Where-Object { $_ -eq $Char }).Count
}

$Max = 0
$Min = [Math]::Pow(10,5)
foreach ($Letter in $Occurrences.Keys) {
    if ($Occurrences[$Letter] -lt $Min) {$Min = $Occurrences[$Letter]}
    if ($Occurrences[$Letter] -gt $Max) {$Max = $Occurrences[$Letter]}
}
Write-Host "Difference after $Steps steps: $($Max - $Min)"