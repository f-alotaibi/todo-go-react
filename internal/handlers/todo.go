package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo_go_react/internal/models"
	"todo_go_react/internal/storage"
)

type TodosHandler struct {
	store storage.Storage
}

func NewTodoHandler(store storage.Storage) *TodosHandler {
	return &TodosHandler{store: store}
}

func (h *TodosHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.store.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func (h *TodosHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoTask
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = h.store.Create(todo)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TodosHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idValue := r.PathValue("id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(w, "Failed to find id", http.StatusBadRequest)
		return
	}
	var todo models.TodoTask
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil || todo.ID != uint(id) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ok := h.store.Update(todo)
	if !ok {
		http.Error(w, "Failed to find id", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TodosHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idValue := r.PathValue("id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(w, "Failed to find id", http.StatusBadRequest)
		return
	}
	ok := h.store.Delete(uint(id))
	if !ok {
		http.Error(w, "Failed to find id", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
