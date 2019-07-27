@echo off

::windows 64
set GOOS=windows
set GOARCH=amd64
go build -o .\build\skinseek-windows-amd64.exe .\skinseek.go

::windows 32
set GOOS=windows
set GOARCH=386
go build -o .\build\skinseek-windows-386.exe .\skinseek.go

::linux 64
set GOOS=linux
set GOARCH=amd64
go build -o .\build\skinseek-linux-amd64 .\skinseek.go

::linux 32
set GOOS=linux
set GOARCH=386
go build -o .\build\skinseek-linux-386 .\skinseek.go

::linux arm5
set GOOS=linux
set GOARCH=arm
set GOARM=5
go build -o .\build\skinseek-linux-arm5 .\skinseek.go

::linux arm6
set GOOS=linux
set GOARCH=arm
set GOARM=6
go build -o .\build\skinseek-linux-arm6 .\skinseek.go

::linux arm7
set GOOS=linux
set GOARCH=arm
set GOARM=7
go build -o .\build\skinseek-linux-arm7 .\skinseek.go

::linux arm64
set GOOS=linux
set GOARCH=arm64
go build -o .\build\skinseek-linux-arm64 .\skinseek.go

::linux mips
set GOOS=linux
set GOARCH=mips
go build -o .\build\skinseek-linux-mips .\skinseek.go

::linux mips64
set GOOS=linux
set GOARCH=mips64
go build -o .\build\skinseek-linux-mips64 .\skinseek.go

::linux mips64le
set GOOS=linux
set GOARCH=mips64le
go build -o .\build\skinseek-linux-mips64le .\skinseek.go

::linux mipsle
set GOOS=linux
set GOARCH=mipsle
go build -o .\build\skinseek-linux-mipsle .\skinseek.go

::linux ppc64
set GOOS=linux
set GOARCH=ppc64
go build -o .\build\skinseek-linux-ppc64 .\skinseek.go

::linux ppc64le
set GOOS=linux
set GOARCH=ppc64le
go build -o .\build\skinseek-linux-ppc64le .\skinseek.go

::darwin amd64
set GOOS=darwin
set GOARCH=amd64
go build -o .\build\skinseek-darwin-amd64 .\skinseek.go

pause