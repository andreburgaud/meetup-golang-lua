# Go, meet Lua

Code supporting the Golang Meetup presentation about calling Go from Lua or embedding Lua in Go.

The slides are available at https://www.slideshare.net/AndreBurgaud/go-meet-lua/AndreBurgaud/go-meet-lua.

## Requirements

* A recent version of Lua
* LuaJIT
* Go (golang)
* Linux, Mac or WSL environment (tested with WSL)
* Optional: UPX, rlwrap

### Lua

* https://www.lua.org/download.html

```
$ curl -R -O http://www.lua.org/ftp/lua-5.3.5.tar.gz
$ tar zxf lua-5.3.5.tar.gz
$ cd lua-5.3.5
$ make linux test
```

### LuaJIT

* https://luajit.org/download.html
*

```
$ curl -R -0 https://luajit.org/download/LuaJIT-2.1.0-beta3.tar.gz
$ tar zxf LuaJIT-2.1.0-beta3.tar.gz
$ cd LuaJIT-2.1.0-beta3
$ make linux test
```

On Mac, if you encounter an error while compiling LuaJIT set the MACOSX environment variable and the number of cores (`-j8`) accordingly. For example:

```
$  make MACOSX_DEPLOYMENT_TARGET=10.14 DEFAULT_CC=clang -j8
```

### Go

Download the most recent version for your operating system at the following page: https://golang.org/dl/

### UPX, rlwrap

If you want to compress the libraries with `UPX`, `make dist` in the example folders, you can find this utility at https://upx.github.io/.

`rlwrap` wraps an REPL with readline facilities, allowing to recall the previous command and to easily modify any expression.

`UPX` and `rlwrap` are also available via the usual package managers like `Homebrew` on Mac, `apt-get` on Debian-based distros, `pacman` on Arch-based distros...

## Development

Clone the repository containing the code examples:

```
$ git clone https://github.com/andreburgaud/meetup-golang-lua
```

### Global Build

You can execute all builds by executing `make` at the root of the project:

```
$ make
```

**Note**: the sub-projects `gco_lua/dofile` and `cgo_lua/repl` will need a copy of the following lua files in their folders in order to compile successfully:

**Lua header files**:

* `lauxlib.h`
* `lua.h`
* `luaconf.h`
* `lualib.h`

**Lua Library file**:

* `liblua.a` (this file is generated when you compile lua. refer to the section **Lua** and **LuaJIT** above to read the deatils about performing the compilation)

You can clean all the subproject by executing `make clean` at the root of the project:

```
$ make clean
```

### Sub-Projects

Each folder containing code examples has a `Makefile` and a `README` with some description and requirements. Some examples, in particular those depending on [CGO](https://golang.org/cmd/cgo/), require to have access to the header files of the Lua source code. The header files are already in the repo, but if you experienced any issue it is better to copy the header files (`*.h`) or static libraries (`*.a`) corresponding to the version of Lua you installed on your system.

In directories containing a `Makefile`, you can execute:

Build:

```
$ make
```

Test (run):

```
$ make test
```

Distribution (compressed packages), dependecy on upx:

```
$ make dist
```

Clean:

```
$ make clean
```

## Licenses

* The code in this repository is available under the [MIT License](LICENSE.md)
* The documentation is released under the [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/) license
