#  Script name:    compileJS.ps1
#  Created on:     2012-06-26 
#  Purpose:        Compilation par closure du JS du site.

$exec = "python $env:gopath/lib/closure/bin/build/closurebuilder.py";
$exec = "$exec --root=$env:gopath/lib/"
$exec = "$exec --root=$env:gopath/src/studash"
$exec = "$exec --output_mode=compiled"
$exec = "$exec --namespace='studash'"
$exec = "$exec --compiler_jar=compiler.jar"
$exec = "$exec --compiler_flags='--compilation_level=ADVANCED_OPTIMIZATIONS'"
$exec = "$exec --output_file='out.js'"


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