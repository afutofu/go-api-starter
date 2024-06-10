package storage

import (
	"sync"

	"github.com/afutofu/go-api-starter/models"
	log "github.com/sirupsen/logrus"
)

var (
	todos      []models.Todo
	nextTodoID int
	users      = make(map[int]models.User)
	userByName = make(map[string]int)
	nextUserID int
	mu         sync.RWMutex
)

/*
SaveUser saves a user's credentials.

Parameters:
  - user: *models.User containing the user data.

Returns:
  - void: The function modifies the in-memory storage directly.
*/
func SaveUser(user *models.User) {
	mu.Lock()
	defer mu.Unlock()
	user.ID = nextUserID
	nextUserID++
	users[user.ID] = *user
	userByName[user.Username] = user.ID
	log.Info("User saved: ", *user)
}

/*
GetUserByName retrieves a user's credentials by username.

Parameters:
  - username: The username of the user.

Returns:
  - models.User: The user's credentials.
  - bool: Indicates whether the user was found.
*/
func GetUserByID(username string) (models.User, bool) {
	mu.RLock()
	defer mu.RUnlock()
	id, exists := userByName[username]
	if !exists {
		return models.User{}, false
	}
	user, found := users[id]
	return user, found
}

/*
GetUserByName retrieves a user's credentials by username.

Parameters:
  - username: The username of the user.

Returns:
  - models.User: The user's credentials.
  - bool: Indicates whether the user was found.
*/
func GetUserByName(username string) (models.User, bool) {
	mu.RLock()
	defer mu.RUnlock()
	id, exists := userByName[username]
	if !exists {
		return models.User{}, false
	}
	user, found := users[id]
	return user, found
}

/*
AddTodo adds a new todo item to the storage.

Parameters:
  - todo: *models.Todo to be added.

Returns:
  - void: The function modifies the in-memory storage directly.
*/
func AddTodo(todo *models.Todo) {
	mu.Lock()
	defer mu.Unlock()
	todo.ID = nextTodoID
	nextTodoID++
	todos = append(todos, *todo)
	log.Info("Todo added: ", *todo)
}

/*
GetTodos retrieves all todo items from the storage.

Returns:
  - []models.Todo: A slice of all todo items.
*/
func GetTodos() []models.Todo {
	mu.RLock()
	defer mu.RUnlock()
	return todos
}

/*
GetTodoByID retrieves a todo item by its ID.

Parameters:
  - id: int representing the ID of the todo item.

Returns:
  - models.Todo: The todo item with the given ID.
  - bool: Indicates whether the todo item was found.
*/
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

/*
UpdateTodoByID updates a todo item by its ID.

Parameters:
  - id: int representing the ID of the todo item to be updated.
  - updatedTodo: *models.Todo containing the updated data.

Returns:
  - bool: Indicates whether the todo item was updated.
*/
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

/*
DeleteTodoByID deletes a todo item by its ID.

Parameters:
  - id: int representing the ID of the todo item to be deleted.

Returns:
  - bool: Indicates whether the todo item was deleted.
*/
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

/*
ClearUsers clears all users and resets the ID counter.
*/
func ClearUsers() {
	users = make(map[int]models.User)
	userByName = make(map[string]int)
	nextUserID = 0
}

/*
ClearTodos clears all todos and resets the ID counter.
*/
func ClearTodos() {
	mu.Lock()
	defer mu.Unlock()
	todos = nil
	nextTodoID = 0
	log.Info("Todos cleared")
}
