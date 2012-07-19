#########################################################################
#  Script permettant de faire la compilation du Javascript sur le site. #
##########################################################################

$namespace = "studash.Student", "studash"
$outfile = "client/_"


# --- Ne pas toucher, variable autogénérés.
$goPath = $env:gopath
$closurePath = $env:closure

$root = "client/js/", "$closurePath/library/"

$buildScript = "$closurePath/library/closure/bin/build/closurebuilder.py"
$closureCompiler = "$closurePath/compiler/build/compiler.jar"

$opt = "--output_mode=compiled --compiler_jar=$closureCompiler "
#$opt = "--output_mode=compiled --compiler_jar=$closureCompiler --compiler_flags='--compilation_level=ADVANCED_OPTIMIZATIONS'"

# --- Batissons la commande.
$exec = "python $buildScript"

foreach ($namespc in $namespace) { $exec = "$exec --namespace='$namespc'"}
foreach ($src in $root) { $exec = "$exec --root='$src'"}
$exec = "$exec $opt --output_file=$outfile"

# --- Executon la commande
#Write-Output $exec
Invoke-Expression $exec