local ffi = require "ffi"
local golib = ffi.load("./golib.so")

ffi.cdef[[
extern char* Pwd();
]]

function pwd()
  return ffi.string(golib.Pwd())
end

print("Current working directory is: " .. pwd())
