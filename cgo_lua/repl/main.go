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

// lua.h: #define lua_pop(L,n)		lua_settop(L, -(n)-1)
func luaPop(L *C.lua_State, i C.int) {
	C.lua_settop(L, -(i)-1)
}

// lua.h:
// #define lua_tostring(L,i)	lua_tolstring(L, (i), NULL)
// LUA_API const char     *(lua_tolstring) (lua_State *L, int idx, size_t *len);
func luaToString(L *C.lua_State, i int) *C.char {
	cstr := C.lua_tolstring(L, (C.int(i)), nil)
	return cstr
}

// lua.h:
// LUA_API int   (lua_pcallk) (lua_State *L, int nargs, int nresults, int errfunc,
//                             lua_KContext ctx, lua_KFunction k);
// #define lua_pcall(L,n,r,f)	lua_pcallk(L, (n), (r), (f), 0, NULL)
func luaPCall(L *C.lua_State, n, r, f C.int) int {
	res := C.lua_pcallk(L, n, r, f, 0, nil)
	return int(res)
}

func handleError(L *C.lua_State) {
	msg := C.lua_tolstring(L, (-1), nil) // Cleanup done by luaPop
	fmt.Fprintf(os.Stderr, "%s\n", C.GoString(msg))
	luaPop(L, 1)
}

func help() {
	msg := `.exit    Exit the repl
.quit    Exit the repl
.help    Print this help
.version Print Lua version
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

// #define LUA_VERSION	"Lua " LUA_VERSION_MAJOR "." LUA_VERSION_MINOR
// #define LUA_VERSION_NUM		503
// LUA_API const lua_Number *(lua_version) (lua_State *L);
func version(L *C.lua_State) {
	verNum := C.lua_version(L)
	C.lua_getglobal(L, C.CString("_VERSION"))
	version := luaToString(L, -1)
	fmt.Printf("%s (%.0f)\n", C.GoString(version), C.double(*verNum))
	luaPop(L, 1)
}

func execute(L *C.lua_State, line string) {
	expression := C.CString(strings.TrimSpace(line))
	defer C.free(unsafe.Pointer(expression))

	if err := C.luaL_loadstring(L, expression); err != 0 {
		handleError(L)
		return
	}

	if err := luaPCall(L, 0, 0, 0); err != 0 {
		fmt.Println(line)
		handleError(L)
		return
	}

	res := luaToString(L, -1)

	msg := C.GoString(res)
	fmt.Print(msg)
	return
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
		case ".version":
			version(L)
			continue
		case ".exit":
			fallthrough
		case ".quit":
			return
		}

		execute(L, line)
	}
}

func main() {
	fmt.Println("Basic Lua REPL in Go!")
	help()
	prompt := ">>> "
	repl(prompt, os.Stdin, os.Stdout)
}
