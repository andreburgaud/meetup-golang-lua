build: 
	go build -o golib.so -buildmode=c-shared main.go
	gcc main.c ./golib.so

dist:
	go build -o golib.so -ldflags="-s -w" -buildmode=c-shared main.go 
	gcc -O3 main.c ./golib.so
	strip a.out
	upx golib.so

test:
	./a.out

clean:
	rm a.out golib.so golib.h || true
