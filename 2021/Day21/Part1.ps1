$Data = Get-Content "$PSScriptRoot/Input.txt"

$Players = @{}

foreach ($Row in $Data) {
    $null = $Row -match 'Player (?<Player>\d+) starting position: (?<Pos>\d+)'
    $Players[$matches['Player']] = @{'Score'=0;'Position'=[Int32]$matches['Pos']}
}

$Winner = $false
$DiceNo = 1
$DiceRolls = 0
$BoardMax = 10
$DiceMax = 100

while (-not $Winner) {

    foreach ($Player in '1','2') {

        # Work out how many spaces the player moves forward
        $Spaces = 0
        for ($i=0; $i -lt 3; $i++) {
            $Spaces += $DiceNo
            if ($DiceNo -eq $DiceMax) {$DiceNo = 1} else {$DiceNo++}
        }
        # Move the player forward this many positions
        $Players[$Player]['Position'] += $Spaces
        # Account for the cyclical board
        while ($Players[$Player]['Position'] -gt $BoardMax) { $Players[$Player]['Position'] -= 10 }

        $Players[$Player]['Score'] += $Players[$Player]['Position']

        $DiceRolls += 3

        if ($Players[$Player]['Score'] -ge 1000) {$Winner = $true; break}
    }

}

switch ($Player) {
    '1' { $LosingScore = $Players['2']['Score'] }
    '2' { $LosingScore = $Players['1']['Score'] }
}

Write-Host "Losing Score * Dice Rolls = $($DiceRolls * $LosingScore)"