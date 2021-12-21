$Data = Get-Content "$PSScriptRoot/Input.txt"

$null = $data -match 'x\=(?<xmin>\d+)\.\.(?<xmax>\d+)\, y\=\-(?<ymax>\d+)\.\.\-(?<ymin>\d+)'

$TargetY = (-1 * $matches['ymin'])..(-1 * $matches['ymax'])

# The max y velocity will be whatever gets to the
#  bottom of the target in one step
# Therefore the starting velocity should be one less than this value

$MaxYVelocity = [Math]::Abs(($TargetY | Measure-Object -Minimum).Minimum) - 1

# Then the max y is the sum of Vy, Vy-1, Vy-2 etc.
$MaxY = 0
for ($i = 0; $i -lt $MaxYVelocity; $i++) { $MaxY += ($MaxYVelocity - $i) }

Write-Host "Max height: $MaxY"