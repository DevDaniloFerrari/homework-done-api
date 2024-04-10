package routes

import "github.com/gorilla/mux"

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", PostTask).Methods("POST")
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
}
