package server

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/74th/vscode-book-golang/model/tasks"
	"github.com/74th/vscode-book-golang/repository"
)

// Server サーバAPI
type Server struct {
	host    string
	webroot string
	server  http.Server
	rep     tasks.Repository
}

// New サーバAPIインスタンスを作成する
func New(addr string, webroot string) *Server {
	s := &Server{
		server: http.Server{
			Addr: addr,
		},
		webroot: webroot,
		rep:     repository.New(),
	}

	s.setRouter()

	return s
}

// Serve サーバを開始する
func (s *Server) Serve() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start server: %s", err.Error())
	}
}

func (s *Server) setRouter() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/tasks", s.list)
	api.POST("/tasks", s.create)
	api.POST("/tasks/:id/done", s.done)

	router.StaticFile("/", filepath.Join(s.webroot, "index.html"))
	router.Static("/js", filepath.Join(s.webroot, "js"))
	s.server.Handler = router
}

// list GET /tasks
func (s *Server) list(c *gin.Context) {
	tasks := s.rep.List()
	c.JSON(http.StatusOK, tasks)
}

// create POST /tasks
func (s *Server) create(c *gin.Context) {
	var task tasks.Task
	c.ShouldBindJSON(&task)
	task.ID = s.rep.Add(task)
	c.JSON(200, &task)
}

// done POST /tasks/:id/done
func (s *Server) done(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
	}
	err = s.rep.Done(id)
	if err != nil {
		c.Status(404)
		return
	}
	c.Status(200)
}
