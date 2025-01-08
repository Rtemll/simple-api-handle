package main

import (
	"fmt"
	"net/http"
	"pantela/internal/database"
	"pantela/internal/handlers"
	"pantela/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)
	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/task", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/task/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	// router.HandleFunc("/api/task/{id}", PatchTaskHandler).Methods("PATCH")
	// router.HandleFunc("/api/task/{id}", DeleteTaskHandler).Methods("DELETE")
	fmt.Println("Server listening on port 8080") // Добавлено для лучшей обратной связи
	http.ListenAndServe(":8080", router)
}
