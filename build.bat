@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go env
go build
pause