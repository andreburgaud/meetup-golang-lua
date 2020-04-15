package main

import (
	"C"
	"os"
)

//export Pwd
func Pwd() *C.char {
	d, err := os.Getwd()
	if err != nil {
		return C.CString("Error")
	}
	return C.CString(d)
}

func main() {}
