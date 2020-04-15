build:
	go build -o golib.so -buildmode=c-shared main.go

dist:
	go build -o golib.so -buildmode=c-shared -ldflags="-s -w" main.go
	upx golib.so

test:
	@luajit getcwd.lua

repl:
	@rlwrap -pgreen luajit