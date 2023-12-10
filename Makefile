cli:
	go build cmd/q/main.go

web:
	GOOS=js GOARCH=wasm tinygo build -o main.wasm cmd/web/main.go
