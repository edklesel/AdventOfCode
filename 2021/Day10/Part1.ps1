$Data = Get-Content "$PSScriptRoot/Input.txt"

$Points = @{')'=3;']'=57;'}'=1197;'>'=25137}
$Opening = '([{<'
$Closing = ')]}>'
$Pairs = @{'('=')';'{'='}';'['=']';'<'='>'}
$Total = 0


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

}

Write-Host "Total points for illegal characters: $Total"