package main

// #cgo LDFLAGS: ${SRCDIR}/liblua.a -lm -ldl
// #include <dlfcn.h>
// #include <stdlib.h>
// #include <stdio.h>
// #include <string.h>
// #include <math.h>
// #include <lauxlib.h>
// #include <lua.h>
// #include <luaconf.h>
// #include <lualib.h>
import "C"

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unsafe"
)

func help() {
	msg := `.exit   Exit the repl
.help   Print this help
-----------------------
Examples:
>>> print "Hello Lua!"
>>> print(_VERSION)
>>> print(type("Lua"))
>>> print(type({}))
>>> print(math.pi)
>>> print(os.date())
>>> print(os.tmpname())
>>> print(os.getenv("HOME"))
>>> print((function (x) return 2*x end)(3))
----------------------`
	fmt.Println(msg)
}

func execute(L *C.lua_State, line string) bool {
	expression := C.CString(strings.TrimSpace(line))
	defer C.free(unsafe.Pointer(expression))

	C.luaL_loadstring(L, expression)

	if err := C.lua_pcallk(L, 0, 0, 0, 0, nil); err != 0 {
		C.lua_settop(L, 0) // Clear stack
		fmt.Println(line)
		return false
	}
	result := C.lua_tolstring(L, (-1), nil)
	defer C.free(unsafe.Pointer(result))

	msg := C.GoString(result)
	fmt.Print(msg)
	return true
}

func repl(prompt string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	L := C.luaL_newstate()
	C.luaL_openlibs(L)
	defer C.lua_close(L)

	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			continue
		}

		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		switch line {
		case ".help":
			help()
			continue
		case ".exit":
			return
		}

		if !execute(L, line) {
			continue
		}
	}
}

func main() {
	fmt.Println("Basic Lua REPL in Go!")
	help()
	prompt := ">>> "
	repl(prompt, os.Stdin, os.Stdout)
}
