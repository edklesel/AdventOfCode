$Data = Get-Content "$PSScriptRoot/Input.txt"

$null = $data -match 'x\=(?<xmin>\d+)\.\.(?<xmax>\d+)\, y\=\-(?<ymax>\d+)\.\.\-(?<ymin>\d+)'

$TargetX = $matches['xmin']..$matches['xmax']
$TargetY = (-1 * $matches['ymin'])..(-1 * $matches['ymax'])

# X and Y are independent, so the values can be measured independently

# From Part 1, the max value Y velocity can be is already calculated
$MaxYVelocity = [Math]::Abs(($TargetY | Measure-Object -Minimum).Minimum) - 1

# Similarly, the smallest y can be is the minimum y value,
#  so it reaches the target in one step
$MinYVelocity = ($TargetY | Measure-Object -Minimum).Minimum

# In a similar vein, the max that the X velcity can be is
#  such that it reaches the target in one step.
$MaxXVelocity = ($TargetX | Measure-Object -Maximum).Maximum

# We can find the minimum by calculating the minimum n where n(n+1)/2
#  falls within the target
$MinXVelocity = 0
while ($TargetX -notcontains ($MinXVelocity * ($MinXVelocity + 1)/2)) { $MinXVelocity++ }

# Now brute force the values for which the probe hits the target
# It takes a while
$Velocities = [System.Collections.Generic.List[String]]@()
foreach ($Vy in $MinYVelocity..$MaxYVelocity) {
    foreach ($Vx in $MinXVelocity..$MaxXVelocity) {

        $x = 0
        $y = 0
        $Velocity = $Vx,$Vy

        while ($true) {

            # If the probe hits the target
            if ($TargetX -contains $x -and $TargetY -contains $y) {
                $Velocities.Add("($Vx,$Vy)")
                break
            }
            # Otherwise if the probe has fallen below, or gone past, the target
            elseif ($y -lt ($TargetY | Measure-Object -Minimum).Minimum) { break }
            elseif ($x -gt ($TargetX | Measure-Object -Maximum).Maximum) { break }

            # Move the probe to the next step
            $x += $Velocity[0]
            $y += $Velocity[1]

            # Change the velocities
            if ($Velocity[0] -ne 0) { $Velocity[0]-- }
            $Velocity[1]--

        }

    }
}