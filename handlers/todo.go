package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/afutofu/go-api-starter/models"
	"github.com/afutofu/go-api-starter/storage"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Error("Failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Creating todo: ", todo)

	storage.AddTodo(&todo)
	log.Info("Todo created: ", todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := storage.GetTodos()
	log.Info("Fetched todos: ", todos)
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Error("Invalid ID: ", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, found := storage.GetTodoByID(id)
	if !found {
		log.Warn("Todo not found: ", id)
		http.NotFound(w, r)
		return
	}

	log.Info("Fetched todo: ", todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Error("Invalid ID: ", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		log.Error("Failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo.ID = id
	if updated := storage.UpdateTodoByID(id, &updatedTodo); !updated {
		log.Warn("Todo not found for update: ", id)
		http.NotFound(w, r)
		return
	}

	log.Info("Todo updated: ", updatedTodo)
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Error("Invalid ID: ", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if deleted := storage.DeleteTodoByID(id); !deleted {
		log.Warn("Todo not found for deletion: ", id)
		http.NotFound(w, r)
		return
	}

	log.Info("Todo deleted: ", id)
	w.WriteHeader(http.StatusNoContent)
}
