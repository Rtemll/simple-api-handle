package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
var task string
func HelloHandler(w http.ResponseWriter, r *http.Request){
fmt.Fprint(w, "Hello, "+task+"!")
}
func PostTaskHandler(w http.ResponseWriter, r *http.Request){
	var newTask struct{
		Task string `json:"task"`
	}
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task = newTask.Task
	fmt.Fprint(w, "Task update successfully")
}
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostTaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}