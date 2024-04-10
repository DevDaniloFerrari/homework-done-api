package routes

import "github.com/gorilla/mux"

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
}
