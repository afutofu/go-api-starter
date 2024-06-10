package storage

import (
	"testing"

	"github.com/afutofu/go-api-starter/models"
	"github.com/stretchr/testify/assert"
)

func TestAddTodoAndGetTodos(t *testing.T) {
	ClearTodos() // Clear the storage before running the test
	todo := models.Todo{Text: "Test Todo", Completed: false}
	AddTodo(&todo)

	todos := GetTodos()
	assert.Len(t, todos, 1, "Expected one todo in the storage")
	assert.Equal(t, "Test Todo", todos[0].Text, "Expected todo title to match")
	assert.Equal(t, false, todos[0].Completed, "Expected todo status to match")
}

func TestGetTodoByID(t *testing.T) {
	ClearTodos() // Clear the storage before running the test
	todo := models.Todo{Text: "Test Todo", Completed: false}
	AddTodo(&todo)

	fetchedTodo, found := GetTodoByID(0)
	assert.True(t, found, "Expected todo to be found")
	assert.Equal(t, "Test Todo", fetchedTodo.Text, "Expected todo title to match")
	assert.Equal(t, false, fetchedTodo.Completed, "Expected todo status to match")
}

func TestUpdateTodoByID(t *testing.T) {
	ClearTodos() // Clear the storage before running the test
	todo := models.Todo{Text: "Test Todo", Completed: false}
	AddTodo(&todo)

	updatedTodo := models.Todo{Text: "Updated Todo", Completed: true}
	updated := UpdateTodoByID(0, &updatedTodo)
	assert.True(t, updated, "Expected todo to be updated")

	fetchedTodo, found := GetTodoByID(0)
	assert.True(t, found, "Expected todo to be found")
	assert.Equal(t, "Updated Todo", fetchedTodo.Text, "Expected updated todo title to match")
	assert.Equal(t, true, fetchedTodo.Completed, "Expected updated todo status to match")
}

func TestDeleteTodoByID(t *testing.T) {
	ClearTodos() // Clear the storage before running the test
	todo := models.Todo{Text: "Test Todo", Completed: false}
	AddTodo(&todo)

	deleted := DeleteTodoByID(0)
	assert.True(t, deleted, "Expected todo to be deleted")

	_, found := GetTodoByID(0)
	assert.False(t, found, "Expected todo to not be found")
}
