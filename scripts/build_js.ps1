#########################################################################
#  Script permettant de faire la compilation du Javascript sur le site. #
##########################################################################

$namespace = "studash"
$outfile = "client/_"


# --- Ne pas toucher, variable autogénérés.
$goPath = $env:gopath
$closurePath = $env:closure

$buildScript = "$closurePath/library/closure/bin/build/closurebuilder.py"
$closureLib = "$closurePath/library/"
$closureCompiler = "$closurePath/compiler/build/compiler.jar"
$srcFolder = "client/js/"

# --- Batissons la commande.
$exec = "python $buildScript --root=$closureLib --root=$srcFolder --output_mode=compiled"
$exec = "$exec --namespace='$namespace' --compiler_jar=$closureCompiler"
$exec = "$exec --compiler_flags='--compilation_level=ADVANCED_OPTIMIZATIONS' --output_file=$outfile"

# --- Executon la commande
#Write-Output $exec
Invoke-Expression $exec