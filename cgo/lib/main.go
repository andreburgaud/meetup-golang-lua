package main

// #include <stdlib.h>
// #include <hello.h>
import "C"

import (
   "unsafe"
)

func main() {
   C.hello()
   name := C.CString("GoMN")
   defer C.free(unsafe.Pointer(name))
   C.print(name)
}
