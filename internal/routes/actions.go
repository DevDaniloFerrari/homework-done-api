package routes

import (
	"encoding/json"
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

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	tasks := service.FindAll()
	json.NewEncoder(writer).Encode(&tasks)
}
