[System.Collections.ArrayList]$Data = Get-Content "$PSScriptRoot/Input.txt"

$Count = 0

$Numbers = @{
    '0' = 'ABCEFG';
    '1' = 'CF';
    '2' = 'ACDEG';
    '3' = 'ACDFG';
    '4' = 'BCDF';
    '5' = 'ABDFG';
    '6' = 'ABDEFG';
    '7' = 'ACF';
    '8' = 'ABCDEFG';
    '9' = 'ABCDFG';
}

<#

0256

1 4 7 and 8 have unique numbers of lines, so these are obvious

3 numbers have 6 lines, 0 6 and 9
3 numbers have 5 lines, 2 3 and 5

4 fits in 9 but no other 6-lined signal, so if a 6-lined
 signal fits 4 then it must be 9.

7 fits in 3 but not 2 or 5, therefore if 7 fits in a
 signal with 5 lines it must be 3.

Both 6 and 9 entirely contain 5,therefore if
 5 doesn't fit in a signal with 6 lines it must be 0



#>

function Compare-Signals {

    param(
        [string]$Signal1,
        [string]$Signal2
    )

    $Array1 = $Signal1 -split '' | Where-Object {$_ -match '\w'}
    $Array2 = $Signal2 -split '' | Where-Object {$_ -match '\w'}

    foreach ($Char in $Array1) { if ($Array2 -notcontains $Char) {return $false} }

    return $true

}

$TotalSum = 0

foreach ($Row in $Data) {

    $Patterns = @{'0'=$null;'1'=$null;'2'=$null;'3'=$null;'4'=$null;'5'=$null;'6'=$null;'7'=$null;'8'=$null;'9'=$null;}

    $Signal = ($Row -split ' \| ' | Select-Object -First 1) -split ' '
    $Output = ($Row -split ' \| ' | Select-Object -Last 1) -split ' '

    # Order the characters first up
    for ($i=0; $i -lt $Signal.Count; $i++) {$Signal[$i] = ($Signal[$i] -split '' | Sort-Object) -join ''}
    for ($i=0; $i -lt $Output.Count; $i++) {$Output[$i] = ($Output[$i] -split '' | Sort-Object) -join ''}

    # Extract the unique numbers
    $Patterns['1'] = ($Signal + $Output) | Where-Object { ($_ | Measure-Object -Character).Characters -eq 2 } | Select-Object -First 1
    $Patterns['4'] = ($Signal + $Output) | Where-Object { ($_ | Measure-Object -Character).Characters -eq 4 } | Select-Object -First 1
    $Patterns['7'] = ($Signal + $Output) | Where-Object { ($_ | Measure-Object -Character).Characters -eq 3 } | Select-Object -First 1
    $Patterns['8'] = ($Signal + $Output) | Where-Object { ($_ | Measure-Object -Character).Characters -eq 7 } | Select-Object -First 1
    
    $5Lines = ($Signal + $Output) | Where-Object { ($_ | Measure-Object -Character | Select-Object -ExpandProperty Characters) -eq 5 } | Sort-Object | Get-Unique
    $6Lines = ($Signal + $Output) | Where-Object { ($_ | Measure-Object -Character | Select-Object -ExpandProperty Characters) -eq 6 } | Sort-Object | Get-Unique

    $Patterns['9'] = $6Lines | Where-Object { Compare-Signals -Signal1 $Patterns['4'] -Signal2 $_ } | Select-Object -First 1
    $Patterns['3'] = $5Lines | Where-Object { Compare-Signals -Signal1 $Patterns['7'] -Signal2 $_ } | Select-Object -First 1

    # Remove the numbers we've found from the 5 and 6 line arrays
    $5Lines = $5Lines | Where-Object { $_ -ne $Patterns['3']}
    $6Lines = $6Lines | Where-Object { $_ -ne $Patterns['9'] }

    # To get 6 and 5 we have to find a which of the remaining 5 lined numbers
    #  fits into which of the remaining 6 lined numbers
    foreach ($5l in $5Lines) {
        foreach ($6l in $6Lines) {
            if (Compare-Signals -Signal1 $5l -Signal2 $6l) { $Patterns['5'] = $5l; $Patterns['6'] = $6l }
        }
    }

    # The remaining 5 and 6 lined numbers must be 0 and 2 respectively
    $Patterns['2'] = $5Lines | Where-Object { $_ -ne $Patterns['5'] } | Select-Object -First 1
    $Patterns['0'] = $6Lines | Where-Object { $_ -ne $Patterns['6'] } | Select-Object -First 1

    $OutputString = ''
    foreach ($Number in $Output) {

        foreach ($PatternID in $Patterns.Keys) {

            if ($Number -eq $Patterns[$PatternID]) {
                $OutputString = "${OutputString}$PatternID"
            }

        }

    }

    $TotalSum += [Int32]$OutputString

}

Write-Host "Sum of output numbers: $TotalSum"