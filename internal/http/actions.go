package http

import (
	"net/http"

	"github.com/DevDaniloFerrari/homeworke-done-api/internal"
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/database"
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/task"
	"github.com/gin-gonic/gin"
)

var service task.Service

func Configure() {
	service = task.Service{
		Repository: task.Repository{
			Conn: database.Conn,
		},
	}
}

func PostTasks(ctx *gin.Context) {
	var task internal.TaskModel
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.Create(task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func GetTasks(ctx *gin.Context) {
	p := service.FindAll()
	ctx.JSON(http.StatusOK, p)
}
