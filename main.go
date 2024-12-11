package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
var task string
func HelloHandler(w http.ResponseWriter, r *http.Request){
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
//  fmt.Fprint(w, "Task added successfully")
jsonData, err := json.Marshal(newTask)
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func main(){
		// Вызываем метод InitDB() из файла db.go
		InitDB()
		// Автоматическая миграция модели Message 
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostTaskHandler).Methods("POST")
	fmt.Println("Server listening on port 8080") // Добавлено для лучшей обратной связи
	http.ListenAndServe(":8080", router)
}