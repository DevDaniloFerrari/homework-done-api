package http

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/tasks", PostTasks)
	g.GET("/tasks", GetTasks)
}
