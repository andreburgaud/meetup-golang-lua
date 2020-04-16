build:
	gofmt -w *.go
	go build -o golua .

dist:
	go build -o golua -ldflags="-s -w" .
	upx golua

test:
	@rlwrap -pred ./golua

clean:
	rm golua || true
