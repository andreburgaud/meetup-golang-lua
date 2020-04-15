# LuaJIT FFI


## REPL

```
$ rlwrap -pgreen luajit
LuaJIT 2.0.5 -- Copyright (C) 2005-2017 Mike Pall. http://luajit.org/
JIT: ON CMOV SSE2 SSE3 SSE4.1 fold cse dce fwd dse narrow loop abc sink fuse
> go = require("go")
> print(go.pwd())
Invoking go code to get the current working directory
/home/andre/AB/git/labs/golang/lua/luajit/ffi
> os.exit()
```

or

```
$ rlwrap -pgreen luajit
LuaJIT 2.0.5 -- Copyright (C) 2005-2017 Mike Pall. http://luajit.org/
JIT: ON CMOV SSE2 SSE3 SSE4.1 fold cse dce fwd dse narrow loop abc sink fuse
> dofile('getcwd.lua')
Invoking go code to get the current working directory
/home/andre/AB/git/labs/golang/lua/luajit/ffi
> os.exit()
```
