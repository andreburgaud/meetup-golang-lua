package main

// #cgo LDFLAGS: ./liblua.a -lm -ldl
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
	"unsafe"
)

const (
	LUA_MULTRET = -1
	NULL = 0
)

func main() {
	L := C.luaL_newstate()
	C.luaL_openlibs(L);
	luaFile := C.CString("hello.lua")
	defer C.free(unsafe.Pointer(luaFile))
	C.luaL_loadfilex(L, luaFile, nil)
	C.lua_pcallk(L, 0, LUA_MULTRET, 0, 0, nil)
	C.lua_close(L)
}

/*
#define luaL_dofile(L, fn) \
	(luaL_loadfile(L, fn) || lua_pcall(L, 0, LUA_MULTRET, 0))

#define luaL_loadfile(L,f)	luaL_loadfilex(L,f,NULL)

#define lua_pcall(L,n,r,f)	lua_pcallk(L, (n), (r), (f), 0, NULL)
*/