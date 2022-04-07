.PHONY: build

goroot:=$(shell go env GOROOT)


build:
	GOOS=js GOARCH=wasm go build -o build/app.wasm github.com/gotrino/fusion-example/cmd/app
	cp "$(goroot)/misc/wasm/wasm_exec.js" build/wasm_exec.js
	cp -R assets/ build/


run: build
	 go run github.com/gotrino/fusion-example/cmd/server
