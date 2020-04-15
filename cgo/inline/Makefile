build: 
	go build -o golua .

dist:
	go build -o golua -ldflags="-s -w" main.go 
	upx golua

test:
	@./golua
