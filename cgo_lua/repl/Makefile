build: 
	go build -o golua .

dist:
	go build -o golua -ldflags="-s -w" . 
	upx golua

test:
	@rlwrap -pred ./golua