package main

import (
	"flag"

	"github.com/74th/devcontainer-book-go/server"
)

func main() {
	var webroot string
	flag.StringVar(&webroot, "w", "./public/html", "web root path")
	flag.Parse()
	sv := server.New("0.0.0.0:8080", webroot)
	sv.Serve()
}
