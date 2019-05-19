package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/74th/vscode-book-golang/repository"
	"github.com/74th/vscode-book-golang/model/tasks"
)

// tasksAPI タスクAPI
type tasksAPI struct {
	rep tasks.Repository
}

// setRouter ルーターの設定
func (a *tasksAPI) setRouter(router *gin.RouterGroup) {
	router.GET("/tasks", a.list)
	router.POST("/tasks", a.create)
	router.POST("/tasks/:id/done", a.done)
}

// init 初期化
func (a *tasksAPI) init(router *gin.RouterGroup) {
	a.rep = repository.New()
	a.setRouter(router)
}

// list GET /tasks
func (a *tasksAPI) list(c *gin.Context) {
	tasks := a.rep.List()
	c.JSON(http.StatusOK, tasks)
}

// create POST /tasks
func (a *tasksAPI) create(c *gin.Context) {
	var task tasks.Task
	c.ShouldBindJSON(&task)
	task.ID = a.rep.Add(task)
	c.JSON(200, &task)
}

// done POST /tasks/:id/done
func (a *tasksAPI) done(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
	}
	err = a.rep.Done(id)
	if err != nil {
		c.Status(404)
		return
	}
	c.Status(200)
}
