@echo off
PUSHD ..\..\
go install studash/adapters
go install studash/errors
go install studash/pages
go install studash/tools
POPD