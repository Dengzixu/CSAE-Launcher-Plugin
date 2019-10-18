@echo off
SET GIN_MODE=release
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
go generate
go build -o "./binary/CSAE_launcher_Plugin_386.exe" Main.go