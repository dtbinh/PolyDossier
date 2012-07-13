################################################
#  Formatte tous les fichiers .go              #
################################################

Get-ChildItem $env:gopath\src\studash | ForEach-Object { go fmt studash/$_ }