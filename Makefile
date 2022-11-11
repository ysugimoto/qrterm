.PHONY: build

build:
	GOOS=darwin go build -o dist/qrterm cmd/main.go

all:
	GOOS=darwin GOARCH=amd64 go build -o dist/qrterm-darwin-amd64 cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/qrterm-darwin-arm64 cmd/main.go

