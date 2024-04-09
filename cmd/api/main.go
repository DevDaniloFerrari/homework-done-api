package main

import (
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "postgresql://posts"
	_, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	g.Run(":4000")
}
