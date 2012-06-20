@echo off
PUSHD ..\%1
FOR %%F IN (*.go) DO go fmt %%F
POPD