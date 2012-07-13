#########################################################################
#  Script permettant de faire la compilation du Javascript sur le site. #
##########################################################################

$namespace = "studash"
$outfile = "result.js"


# Ne pas toucher, variable autogénérés.
$goPath = $env:gopath
$closurePath = $env:closure

$buildScript = "$closurePath/library/closure/bin/build/closurebuilder.py"
$closureLib = "$closurePath/library/"
$closureCompiler = "$closurePath/compiler/build/compiler.jar"
$srcFolder = "$goPath/src/studash"


#  Script name:    compileJS.ps1
#  Created on:     2012-06-26 
#  Purpose:        Compilation par closure du JS du site.

$exec = "python $buildScript --root=$closureLib --root=$srcFolder --output_mode=compiled"
$exec = "$exec --namespace='$namespace' --compiler_jar=$closureCompiler"
$exec = "$exec --compiler_flags='--compilation_level=ADVANCED_OPTIMIZATIONS' --output_file=$outfile"


#Write-Output $exec
Invoke-Expression $exec


#$compile = "java -jar compiler.jar";
#$fileEntries = [IO.Directory]::GetFiles("$PWD\..\client\js"); 
#foreach($fileName in $fileEntries) 
#{
#  $compile = "$compile --js=$fileName";
#}

#$compile = "$compile --js_output_file=out.js"

#Write-Output $compile
#Invoke-Expression $compile

#closure-library/closure/bin/build/closurebuilder.py \
#  --root=closure-library/ \
#  --root=myproject/ \
#  --namespace="myproject.start" \
#  --output_mode=compiled \
#  --compiler_jar=compiler.jar \
#  --compiler_flags="--compilation_level=ADVANCED_OPTIMIZATIONS" \
#  > start-compiled.js