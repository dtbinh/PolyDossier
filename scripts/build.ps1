################################################
#  Script préparant le programme à être lancé. #
################################################


Write-Host "Compiling Go code"
./scripts/build_go.ps1

Write-Host "Compiling JS code"
./scripts/build_js.ps1