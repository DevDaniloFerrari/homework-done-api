package main

import (
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/database"
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "postgresql://postgres:1234@localhost:5432/homework"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.Default()
	http.Configure()
	http.SetRoutes(g)
	g.Run(":4000")
}
