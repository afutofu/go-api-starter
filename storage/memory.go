package storage

import (
	"sync"

	"github.com/afutofu/go-api-starter/models"
	log "github.com/sirupsen/logrus"
)

var (
	todos  []models.Todo
	nextID int
	mu     sync.RWMutex
)

func AddTodo(todo *models.Todo) {
	mu.Lock()
	defer mu.Unlock()
	todo.ID = nextID
	nextID++
	todos = append(todos, *todo)
	log.Info("Todo added: ", *todo)
}

func GetTodos() []models.Todo {
	mu.RLock()
	defer mu.RUnlock()
	return todos
}

func GetTodoByID(id int) (models.Todo, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, todo := range todos {
		if todo.ID == id {
			return todo, true
		}
	}
	return models.Todo{}, false
}

func UpdateTodoByID(id int, updatedTodo *models.Todo) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = *updatedTodo
			log.Info("Todo updated: ", *updatedTodo)
			return true
		}
	}
	return false
}

func DeleteTodoByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			log.Info("Todo deleted: ", id)
			return true
		}
	}
	return false
}

func ClearTodos() {
	mu.Lock()
	defer mu.Unlock()
	todos = nil
	nextID = 0
	log.Info("Todos cleared")
}
