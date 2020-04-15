#include <stdio.h>
#include "hello.h"

void hello() {
    printf("C code (hello) >>> ");
    printf("Hello from C!\n");
}

void print(char* msg) {
    printf("C code (print) >>> ");
    printf("Hello %s!\n", msg);
}
