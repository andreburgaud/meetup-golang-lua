package main

import (
	"github.com/yuin/gopher-lua"
)


func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("Hello GoMN!")`); err != nil {
		panic(err)
	}
}


