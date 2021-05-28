
all: build build-cmd


local:
	go build -mod vendor -tags "jsoniter" -o app_api

build:
	go build -mod vendor -tags "jsoniter" -o bin/app_api

build-cmd:
	go build -mod vendor -tags "jsoniter" -o bin/ ./cmd/...
