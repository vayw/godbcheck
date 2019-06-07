.PHONY: test
all: build done

build:
	@echo "Building..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-X main.version=$(shell git describe --always --tags) -s -w'
build-upx:
	@echo "Building..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=$(shell git describe --always --tags)"
	upx holzhacker
done:
	@echo "Done."

