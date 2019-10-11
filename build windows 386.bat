@echo on
SET GIN_MODE=release
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
go build
