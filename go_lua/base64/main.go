package main

import (
	"encoding/base64"
	"github.com/yuin/gopher-lua"
)

func base64Encode(l *lua.LState) int {
	msg := l.CheckString(1)
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	l.Push(lua.LString(encoded))
	return 1
}

func register(l *lua.LState) {
	meta := l.NewTypeMetatable("base64")
	l.SetGlobal("base64", meta)
	l.SetField(meta, "encode", l.NewFunction(base64Encode))
}

func main() {
	l := lua.NewState()
	defer l.Close()
	register(l)
	if err := l.DoFile("interact.lua"); err != nil {
		panic(err)
	}
}
