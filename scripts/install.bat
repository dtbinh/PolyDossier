@echo off
PUSHD ..\src
go install studash/adapters
go install studash/errors
go install studash/pages
go install studash/tools

echo Installation de dependences.
cd bin\
go build studash\main
echo Compilation de Main.
POPD