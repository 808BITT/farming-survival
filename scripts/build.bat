@echo off

:: Build the project to bin
go build -o bin\main.exe main.go

:: Run the project
cd bin
main.exe
cd ..