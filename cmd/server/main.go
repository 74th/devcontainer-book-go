package main

import (
	"flag"

	"github.com/74th/vscode-book-golang/server"
)

func main() {
	var webroot string
	flag.StringVar(&webroot, "w", "./public/html", "web root path")
	flag.Parse()
	sv := server.New("127.0.0.1:8080", webroot)
	sv.Serve()
}
