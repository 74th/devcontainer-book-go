package server

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Server サーバAPI
type Server struct {
	host     string
	webroot  string
	server   http.Server
	tasksAPI tasksAPI
}

// New サーバAPIインスタンスを作成する
func New(addr string, webroot string) *Server {
	s := &Server{
		server: http.Server{
			Addr: addr,
		},
		webroot: webroot,
	}

	router := gin.Default()
	s.server.Handler = router

	apiRouter := router.Group("/api")
	s.tasksAPI.init(apiRouter)

	router.StaticFile("/", filepath.Join(webroot, "index.html"))
	router.Static("/js", filepath.Join(webroot, "js"))

	return s
}

// Serve サーバを開始する
func (s *Server) Serve() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start server: %s", err.Error())
	}
}
