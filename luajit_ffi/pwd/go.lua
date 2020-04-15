local go = {}

local ffi = require "ffi"
local lib = ffi.load("./golib.so")

ffi.cdef[[
extern char* Pwd();
]]

go.pwd = function()
  return ffi.string(lib.Pwd())
end

return go
