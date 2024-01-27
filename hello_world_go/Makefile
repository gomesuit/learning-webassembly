build:
	GOOS=js GOARCH=wasm go build -o dist/hello.wasm main.go

copy:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./dist/wasm_exec.js

http:
	open http://localhost:8080
	go run server.go
