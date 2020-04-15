build:
	gcc -c hello.c
	ar rs hello.a *.o 
	go build -o golua .

dist:
	gcc -c hello.c
	ar rs hello.a *.o
	strip hello.a 
	go build -o golua -ldflags="-s -w" .
	upx golua

test:
	@./golua

clean:
	rm *.o *.a golua || true