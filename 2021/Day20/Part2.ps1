$Data = Get-Content "$PSScriptRoot/Input.txt"

$Key = $Data[0]

$Picture = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()
foreach ($Row in $Data[2..($Data.Count)]) { $Picture.Add([System.Collections.Generic.List[String]]($Row -split '' | Where-Object {$_ -in '.','#'})) }

$Steps = 50
$Buffer = 1

# From the key, we need to work out what the behaviour of the pixels outside the picture will be. For the first buffering, the new character will always be .
$Chars = @{'0'='.'}
# As they start as a . if key[0] is . then they will always remain a .
if ($Key[0] -eq '.') {
    $Chars['Odd']='.'
    $Chars['Even']='.'
}
# If key[0] = # then after step 1, they will flip to #
if ($Key[0] -eq '#') {
    $Chars['Odd']='#'
    # Then, if key[511] is . they will flip back
    if ($Key[511] -eq '.') {$Chars['Even'] = '.'}
    # Otherwise, they stay # and the number of pixels is infinite
    else { $Chars['Even'] = '#' }
}
# The area which always flips is the pixels which lie outside a 1 pixel buffer of the starting picture, therefore we only need to care about them

function Format-Picture {

    param([System.Collections.Generic.List[System.Collections.Generic.List[String]]]$Picture)

    $Output = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()

    $local:Buffer = $global:Buffer
    $GridRows = $Picture.Count
    $GridCols = $Picture[0].Count

    if     ($Global:Step -eq 0)     {$Char = $global:Chars['0']}
    elseif ($Global:Step % 2 -eq 1) {$Char = $global:Chars['Odd']}
    elseif ($Global:Step % 2 -eq 0) {$Char = $global:Chars['Even']}

    # Do the top buffer row(s)
    for ($Row = 0; $Row -lt $Buffer; $Row++) {
        $Output.Add([System.Collections.Generic.List[String]]@())
        for ($i=0; $i -lt $GridCols+(2*$Buffer); $i++) { $Output[$Row].Add($Char) }
    }

    for ($Row=0; $Row -lt $GridRows; $Row++) {
        # Add the dot(s) at the start
        $Output.Add([System.Collections.Generic.List[String]]@())
        for ($i=0; $i -lt $Buffer; $i++) {$Output[$Row + $Buffer].Add($Char)}
        # Add the row from the picture
        for ($i=0; $i -lt $GridCols; $i++) {$Output[$Row + $Buffer].Add($Picture[$Row][$i])}
        # Add the dot(s) at the end
        for ($i=0; $i -lt $Buffer; $i++) {$Output[$Row + $Buffer].Add($Char)}
    }

    # Add the final buffer row(s)
    for ($Row = $GridRows+$Buffer; $Row -lt $GridRows + (2*$Buffer); $Row++) {
        $Output.Add([System.Collections.Generic.List[String]]@())
        for ($i=0; $i -lt $GridCols+(2*$Buffer); $i++) { $Output[$Row].Add($Char) }
    }

    return $Output

}

function Show-Picture { $global:Picture | foreach-object {Write-Output ($_ -join '')} }


# Make the cursor the bottom-right
for ($Step=0; $Step -lt $Steps; $Step++) {

    Write-Host "$($Step+1)/$Steps"

    if     ($Global:Step -eq 0)     {$OutsideChar = $global:Chars['0']}
    elseif ($Global:Step % 2 -eq 1) {$OutsideChar = $global:Chars['Odd']}
    elseif ($Global:Step % 2 -eq 0) {$OutsideChar = $global:Chars['Even']}


    $Picture = Format-Picture -Picture $Picture
    $xmax = $Picture[0].Count
    $ymax = $Picture.Count

    $NextPicture = [System.Collections.Generic.List[System.Collections.Generic.List[String]]]@()

    $GenerateImageTime = Measure-Command {
    for ($y = 0; $y -lt $Picture.Count; $y++) {

        $NextPicture.Add([System.Collections.Generic.List[String]]@())

        for ($x = 0; $x -lt $Picture[0].Count; $x++) {

            $BinaryString = ''
            for ($yy=($y-1); $yy -le ($y+1); $yy++) {
                for ($xx = ($x-1); $xx -le ($x+1); $xx++) {

                    # Write-Host $xx,$yy
                    # If xx or yy are outside the picture, the bit is a .
                    if ($xx -in -1,$xmax -or $yy -in -1,$ymax) { $Char = $OutsideChar }
                    # Otherwise, the bit comes from the picture
                    else { $Char = $Picture[$yy][$xx] }

                    $BinaryString = "${BinaryString}$Char"

                }
            }

            $BinaryString = $BinaryString.Replace('.','0')
            $BinaryString = $BinaryString.Replace('#','1')

            $Index = [Convert]::ToInt32($BinaryString,2)

            $NextPicture[$y].Add($Key[$Index])

        }
    }
    }

    Write-Host "$($Picture.count)x$($Picture[0].Count): $($GenerateImageTime.TotalMilliseconds)ms ($($GenerateImageTime.TotalMilliseconds / (($Picture.count)*($Picture[0].Count)))ms per block) "

    $Picture = $NextPicture

}

$Total = 0
for ($i=0; $i -lt $Picture.Count; $i++) { $Total += ($Picture[$i] | Where-Object {$_ -eq '#'}).Count }
Write-Host "Total Lit Pixels: $Total"