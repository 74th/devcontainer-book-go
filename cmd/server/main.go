package main

import (
	"github.com/74th/vscode-book-golang/server"
)

func main() {
	sv := server.New("127.0.0.1:8080", "/Users/nnyn/Documents/vscode-book-typescript/public/html")
	sv.Serve()
}
