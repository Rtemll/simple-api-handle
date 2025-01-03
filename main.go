package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
func GetTasksHandler(w http.ResponseWriter, r *http.Request){
	var messages []Message
	DB.Find(&messages) // Запрос всех сообщений из базы данных
	// Преобразование слайса в JSON
	jsonData, err := json.Marshal(messages)
	if err != nil {
	 http.Error(w, err.Error(), http.StatusInternalServerError)
	 return
	}
   
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request){
	var newTask Message // Используем структуру Message
 err := json.NewDecoder(r.Body).Decode(&newTask)
 if err != nil {
  http.Error(w, err.Error(), http.StatusBadRequest)
  return
 }

 DB.Create(&newTask) // Сохранение в базу данных
jsonData, err := json.Marshal(newTask)
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID, _ := strconv.Atoi(params["id"])
   
	var updatedTask Message
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)
   
	var existingTask Message
	result := DB.First(&existingTask, taskID) //ищем по ID
	   if result.Error != nil {
	 w.WriteHeader(http.StatusNotFound)
		   json.NewEncoder(w).Encode(map[string]string{"message": "Задача не найдена"})
		   return
	   }
   
	if updatedTask.Task != "" {
	 existingTask.Task = updatedTask.Task
	}
	   if updatedTask.IsDone != existingTask.IsDone { //проверка, что значение изменилось
	 existingTask.IsDone = updatedTask.IsDone
	}
	DB.Save(&existingTask) // Сохранение в базу данных
	jsonData, err := json.Marshal(existingTask)
	   if err != nil {
		   http.Error(w, err.Error(), http.StatusInternalServerError)
		   return
	   }
	w.Write(jsonData)
   }

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, _ := strconv.Atoi(params["id"])
   
	var existingTask Message
	   result := DB.First(&existingTask, taskID) //ищем по ID
	   if result.Error != nil {
		   w.WriteHeader(http.StatusNotFound)
		   json.NewEncoder(w).Encode(map[string]string{"message": "Задача не найдена"})
		   return
	   }
	   DB.Delete(&existingTask)
	w.WriteHeader(http.StatusNoContent)
   }

func main(){
		// Вызываем метод InitDB() из файла db.go
		InitDB()
		// Автоматическая миграция модели Message 
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/task", GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/task", PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/task/{id}", PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/task/{id}", DeleteTaskHandler).Methods("DELETE")
	fmt.Println("Server listening on port 8080") // Добавлено для лучшей обратной связи
	http.ListenAndServe(":8080", router)
}