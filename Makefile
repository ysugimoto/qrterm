.PHONY: build

build:
	GOOS=darwin go build -o dist/qrterm cmd/main.go
