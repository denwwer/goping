.PHONY: run build build-arm

BUILD_FLAGS = -ldflags="-s -w" -trimpath

run:
	@go run main.go

build:
	@GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/goping main.go

build-arm:
	@GOOS=linux GOARCH=arm64 go build $(BUILD_FLAGS) -o build/goping main.go

