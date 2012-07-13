################################################
#  Script préparant le programme à être lancé. #
################################################

Get-ChildItem $env:gopath\src\studash | ForEach-Object { go fmt studash/$_; go install studash/$_ }
go build studash/main