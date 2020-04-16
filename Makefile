all: cgo_export cgo_inline cgo_lib cgo_lua_dofile cgo_lua_repl \
go_lua_b64 go_lua_dofile go_lua_dostring

clean: cgo_export_clean cgo_inline_clean cgo_lib_clean cgo_lua_dofile_clean cgo_lua_repl_clean \
go_lua_b64_clean go_lua_dofile_clean go_lua_dostring_clean

cgo_export:
	$(MAKE) -C cgo/export
	$(MAKE) test -C cgo/export

cgo_export_clean:
	$(MAKE) clean -C cgo/export

cgo_inline:
	$(MAKE) -C cgo/inline
	$(MAKE) test -C cgo/inline

cgo_inline_clean:
	$(MAKE) clean -C cgo/inline

cgo_lib:
	$(MAKE) -C cgo/lib
	$(MAKE) test -C cgo/lib

cgo_lib_clean:
	$(MAKE) clean -C cgo/lib

cgo_lua_dofile:
	$(MAKE) -C cgo_lua/dofile
	$(MAKE) test -C cgo_lua/dofile

cgo_lua_dofile_clean:
	$(MAKE) clean -C cgo_lua/dofile

cgo_lua_repl:
	$(MAKE) -C cgo_lua/repl

cgo_lua_repl_clean:
	$(MAKE) clean -C cgo_lua/repl

go_lua_b64:
	$(MAKE) -C go_lua/base64

go_lua_b64_clean:
	$(MAKE) clean -C go_lua/base64

go_lua_dofile:
	$(MAKE) -C go_lua/dofile

go_lua_dofile_clean:
	$(MAKE) clean -C go_lua/dofile

go_lua_dostring:
	$(MAKE) -C go_lua/dostring

go_lua_dostring_clean:
	$(MAKE) clean -C go_lua/dostring
