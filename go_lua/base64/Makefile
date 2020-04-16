build:
	go build -o golua .

# Only execute mod the first time or to refresh the project
mod:
	go mod init golua

dist:
	go build -o golua -ldflags="-s -w" .
	upx golua

test:
	@./golua

clean:
	rm golua || true