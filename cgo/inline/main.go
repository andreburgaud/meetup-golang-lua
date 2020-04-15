package main

/*
#include <stdio.h>
void hello() {
    printf("Hello from C\n");
}
*/
import "C" // No new line between the block comment and 'import "C"'

func main() {
    C.hello()
}
