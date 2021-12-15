$Data = Get-Content "$PSScriptRoot/Input.txt"

$Points = @{')'=3;']'=57;'}'=1197;'>'=25137}
$Pt2Points = @{')'=1;']'=2;'}'=3;'>'=4}
$Opening = '([{<'
$Closing = ')]}>'
$Pairs = @{'('=')';'{'='}';'['=']';'<'='>'}
$Total = 0
$Pt2Total = [System.Collections.Generic.List[Int64]]@()


for ($i=0; $i -lt $Data.Count; $i++) {

    $Break = $false
    $Row = $Data[$i] -split '' | Where-Object {$_ -ne '' } | ForEach-Object {[String]$_}
    $Expected = [System.Collections.ArrayList]@()
    $j = 0

    foreach ($Char in $Row) {

        # If the character is an opener, log the next expected closer
        if ($Opening -match [Regex]::Escape($Char)) {
            $Expected.Add($Pairs[$Char]) | Out-Null
        }

        # Otherwise if it's a closer
        elseif ($Closing -match [Regex]::Escape($Char)) {

            # If the character isn't the next expected closer, increase the score
            if ($Char -ne $Expected[-1]) {
                $Total += $Points[$Char]
                $Break = $true
            }

            # Otherwise if it is the expected one, remove it from the expected array and carry on
            elseif ($Char -eq $Expected[-1]) {
                $Expected.RemoveAt($Expected.Count - 1)
            }
        }  

        $j++

        if ($Break) {break}

    }

    if (!$Break) {

        $Pt2Total_Row = 0

        for ($j = $Expected.Count -1; $j -ge 0; $j --) {

            $Pt2Total_Row = $Pt2Total_Row * 5

            $Pt2Total_Row += $Pt2Points[$Expected[$j]]



        }

        $Pt2Total.Add($Pt2Total_Row)

    }

}

Write-Host "Total points for illegal characters: $Total"

$Pt2Total = $Pt2Total | Sort-Object

Write-Host "Middle score for incomplete characters: $($Pt2Total[($Pt2Total.Count + 1)/2 -1])"