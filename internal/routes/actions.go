package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/DevDaniloFerrari/homeworke-done-api/internal"
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/database"
	"github.com/DevDaniloFerrari/homeworke-done-api/internal/task"
)

var service task.Service

func Configure() {
	service = task.Service{
		Repository: task.Repository{
			Conn: database.Conn,
		},
	}
}

func PostTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var task internal.TaskModel

	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Create(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the inserted task in the response
	json.NewEncoder(writer).Encode(task)
}

func UpdateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var task internal.TaskModel

	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Update(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated task in the response
	json.NewEncoder(writer).Encode(task)
}

func DeleteTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Extract the task ID from the URL query parameters
	taskIDStr := request.URL.Query().Get("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Delete the task from the database
	err = service.Delete(taskID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Task deleted successfully"))
}

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	tasks := service.FindAll()
	json.NewEncoder(writer).Encode(&tasks)
}
