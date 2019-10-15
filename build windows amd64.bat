@echo off
SET GIN_MODE=release
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
go generate
go build -o "./binary/CSAE_launcher_Plugin_amd64.exe"