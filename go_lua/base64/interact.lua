-- interact.lua

io.write("Enter a string to base64 encode: ")
local msg = io.read()
encoded = base64.encode(msg)
print("The base64 encoded value of '" .. msg .. "'' is '" .. encoded .. "'")