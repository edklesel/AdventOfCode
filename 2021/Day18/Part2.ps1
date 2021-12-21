$Data = Get-Content "$PSScriptRoot/Input.txt"

function Split-Number {
    param([String]$Number)

    $Result = [System.Collections.Generic.List[String]]@()

    $i = 0
    while ($true) {

        if ($Number[$i] -match '\d' -and $Number[$i+1] -match '\d') { $Result.Add("$($Number[$i])$($Number[$i+1])"); $i+=2 }
        else { $Result.Add($Number[$i]); $i++ }

        if ($i -eq $Number.Length) { break } 
    }

    return $Result

}

function Get-Magnitude {
    param([String]$Number)

    $Split = Split-Number -Number $Number

    # Work out the middle comman of this number
    $Brackets = 0
    for ($i=1; $i -lt $Split.Count; $i++) {
        switch ($Split[$i]) {
            '[' { $Brackets++ }
            ']' { $Brackets-- }
        }

        if ($Brackets -eq 0 -and $Split[$i] -eq ',') { break }

    }

    $Left = $Split[1..($i - 1)] -join ''
    $Right = $Split[($i + 1)..($Split.Count - 2)] -join ''

    while ($Left -notmatch '^\d+$') {
        $Left = Get-Magnitude -Number $Left }
    while ($Right -notmatch '^\d+$') {
        $Right = Get-Magnitude -Number $Right }
    
    return (3*[Int]$Left + 2*[Int]$Right)

}

function Reduce-Number {

    param([String]$Number)

    while ($true) {

        # Check for exploding pairs
        $OpenBrackets = 0
        $NeedsExploding = $false
        $NeedsSplitting = $false
        for ($j=0; $j -lt $Number.Length; $j++) {
            switch ($Number[$j]) {
                '[' { $OpenBrackets++ }
                ']' { $OpenBrackets-- }
            }

            if ($OpenBrackets -eq 5 -and $NeedsExploding -eq $false -and (($Number[$j..($j+6)] -join '') -match '^\[\d+\,\d+\]')) { $NeedsExploding = $true; $ExplodeIndex = $j }

            elseif ($Number[$j] -match '\d' -and $NeedsSplitting -eq $false -and $Number[$j+1] -match '\d') { $NeedsSplitting = $true; $SplitIndex = $j }

        }
        # If there is a number within nested pairs
        if ($NeedsExploding) {

            $j = $ExplodeIndex

            $Number_PreExplode = $Number
            switch ($Number[$j+2] -match '\d') {
                $false {
                    $Number_L = [Int]([String]($Number_PreExplode[$j+1]))
                    switch ($Number[$j+4] -match '\d') {
                        $true  {
                            $RightSearch = $j + 5
                            $Number_R = [Int]([String]($Number_PreExplode[($j+3)..($j+4)] -join ''))
                        }
                        $false {
                            $RightSearch = $j + 4
                            $Number_R = [Int]([String]($Number_PreExplode[$j+3]))
                        }
                    }
                }
                $true {
                    $Number_L = [Int]([String]($Number_PreExplode[($j+1)..($j+2)] -join ''))
                    switch ($Number[$j+5] -match '\d') {
                        $true  {
                            $RightSearch = $j+6
                            $Number_R = [Int]([String]($Number_PreExplode[($j+4)..($j+5)] -join ''))
                        }
                        $false {
                            $RightSearch = $j+5
                            $Number_R = [Int]([String]($Number_PreExplode[$j+4]))
                        }
                    }
                }
            }

            $Number_PostExplode_L = Split-Number -Number ($Number_PreExplode[0..($j-1)] -join '')
            $Number_PostExplode_R = Split-Number -Number ($Number_PreExplode[($RightSearch+1)..$($Number_PreExplode.Length -1)] -join '')
            
            # Find the previous number in the number
            for ($ll=($Number_PostExplode_L.Count -1); $ll -ge 0; $ll--) {
                if ($Number_PostExplode_L[$ll] -match '\d+') {
                    $Number_PostExplode_L[$ll] = $Number_L + [Int32]$Number_PostExplode_L[$ll]
                    break
                }
            }
            # Find the next number in the character
            for ($rr=0; $rr -lt ($Number_PostExplode_R.Count); $rr++) {
                if ($Number_PostExplode_R[$rr] -match '\d+') {
                    $Number_PostExplode_R[$rr] = $Number_R + [Int32]$Number_PostExplode_R[$rr]
                    break
                }
            }
            $Number_PostExplode_L = $Number_PostExplode_L -join ''
            $Number_PostExplode_R = $Number_PostExplode_R -join ''

            $Number = "${Number_PostExplode_L}0${Number_PostExplode_R}"

        }
        # Otherwise, if there is a number 10 or larger (or 2 digits)
        elseif ($NeedsSplitting) {

            $j = $SplitIndex
            
            $Number_PreSplit = $Number
            
            $Number_PostSplit_L = $Number_PreSplit[0..($j-1)] -join ''
            $Number_PostSplit_R = $Number_PreSplit[($j+2)..($Number_PreSplit.Length -1)] -join ''
        
            $SplitNumber = [Int]([String]"$($Number_PreSplit[$j])$($Number_PreSplit[$j+1])")

            switch ($SplitNumber % 2) {
                1 {
                    [Int]$SplitNumber_L = ([Decimal]$SplitNumber)/2 -0.5
                    [Int]$SplitNumber_R = ([Decimal]$SplitNumber)/2 +0.5
                }
                0 {
                    $SplitNumber_L = $SplitNumber / 2
                    $SplitNumber_R = $SplitNumber / 2
                }
            }

            $Number = "${Number_PostSplit_L}[${SplitNumber_L},${SplitNumber_R}]${Number_PostSplit_R}"

        }
        else {
            break
        }

    }

    return $Number

}

$BiggestMagnitude = 0
for ($i=0; $i -lt $Data.Count; $i++) {

    Write-Host "$i / $($Data.Count -1)"

    for ($j=0; $j -lt $Data.Count; $j++) {

        if ($i -eq $j) {continue}

        $Number = Reduce-Number -Number "[$($Data[$i]),$($Data[$j])]"
        $Mag = Get-Magnitude -Number $Number

        if ($Mag -gt $BiggestMagnitude) {
            $BiggestMagnitude = $Mag
            $WinningNumbers = @($Data[$i],$Data[$j],$Number)
        }

    }

}

Write-Host "Biggest Magnitude: $BiggestMagnitude"
Write-Host "Winning Numbers:  $($WinningNumbers[0]) + $($WinningNumbers[1]) = $($WinningNumbers[2])"