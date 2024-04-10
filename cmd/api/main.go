package main

import (
	"log"
	"net/http"

	"github.com/DevDaniloFerrari/homeworke-done-api/internal/database"
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	connectionString := "postgresql://postgres:1234@localhost:5432/homework"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := mux.NewRouter()

	routes.Configure()
	routes.SetRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8100"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	errorServer := http.ListenAndServe(":8080", handler)
	if errorServer != nil {
		log.Fatalln("There's an error with the server,", errorServer)
	}
}
