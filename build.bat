set CGO_ENABLED=0
::x86
set GOARCH=386

set GOOS=windows
go build -o ./build/gitup-windows-386.exe

set GOOS=linux
go build -o ./build/gitup-linux-386

set GOOS=freebsd
go build -o ./build/gitup-freebsd-386

::set GOOS=darwin
::go build -o gitup-darwin-386
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::x64
set GOARCH=amd64

set GOOS=windows
go build -o ./build/gitup-windows-amd64.exe

set GOOS=linux
go build -o ./build/gitup-linux-amd64

set GOOS=freebsd
go build -o ./build/gitup-freebsd-amd64

::set GOOS=darwin
::call make.bat --no-clean
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::arm
::set GOARCH=arm
::set GOOS=linux
::call make.bat --no-clean
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
set CGO_ENABLED=
set GOARCH=
set GOOS=