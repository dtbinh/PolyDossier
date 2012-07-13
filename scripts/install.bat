################################################
#  Script préparant le programme à être lancé. #
################################################



cd @echo off
PUSHD ..\..\..\
go install studash/adapters
go install studash/errors
go install studash/pages
go install studash/tools

echo Installation de dependences.
cd bin\
go build studash\main
echo Compilation de Main.
POPD