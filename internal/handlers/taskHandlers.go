package handlers

import (
	"encoding/json"
	"net/http"
	"pantela/internal/taskService"
	"strconv"

	"github.com/gorilla/mux"
)

//структура Handler необходима для внедрения зависимости и упрощения работы с сервисом в дальнейшем
type Handler struct{
	Service *taskService.TaskService
}
// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler{
return &Handler{
	Service: service,
}
}


func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request){
tasks, err := h.Service.GetAllTasks()
if err != nil {
	http.Error(w, err.Error(),http.StatusInternalServerError)
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(tasks)
}
// мы убрали логику работы с БД из нашей ручки, что является хорошей практикой и правильным подходом в целом. Теперь мы просто обращаемся к функции сервиса GetAllTasks, которая обращается в репозиторий, который возвращает все задачи.

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request){
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}


func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	   if err != nil {
		   http.Error(w, "Invalid task ID", http.StatusBadRequest)
		   return
	   }
   
	var task taskService.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
	 http.Error(w, err.Error(), http.StatusBadRequest)
	 return
	}
   
	updatedTask, err := h.Service.UpdateTask(id, task)
	if err != nil {
	 http.Error(w, err.Error(), http.StatusInternalServerError)
	 return
	}
   
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
   }
   
   func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
	 http.Error(w, "Invalid task ID", http.StatusBadRequest)
	 return
	}
   
	err = h.Service.DeleteTask(id)
	if err != nil {
	 http.Error(w, err.Error(), http.StatusInternalServerError)
	 return
	}
   
	w.WriteHeader(http.StatusNoContent)
   }