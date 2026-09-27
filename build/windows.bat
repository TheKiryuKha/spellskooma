@echo off

set CGO_ENABLED=1
go build -tags win -o spellsk00ma.exe .

if %errorlevel% neq 0 (
    echo.
    echo BUILD FAILED
    exit /b %errorlevel%
)

echo.
echo BUILD SUCCESSFUL