package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/afutofu/go-api-starter/models"
	"github.com/afutofu/go-api-starter/storage"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoHandler(t *testing.T) {
	storage.ClearTodos() // Clear the storage before running the test
	todo := models.Todo{Text: "Test Todo", Completed: false}
	body, _ := json.Marshal(todo)

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTodo)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code to be 201")
	var createdTodo models.Todo
	json.Unmarshal(rr.Body.Bytes(), &createdTodo)
	assert.Equal(t, todo.Text, createdTodo.Text, "Expected todo title to match")
	assert.Equal(t, todo.Completed, createdTodo.Completed, "Expected todo status to match")
}

func TestGetTodosHandler(t *testing.T) {
	storage.ClearTodos() // Clear the storage before running the test
	storage.AddTodo(&models.Todo{Text: "Test Todo 1", Completed: false})
	storage.AddTodo(&models.Todo{Text: "Test Todo 2", Completed: false})

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTodos)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200")
	var todos []models.Todo
	json.Unmarshal(rr.Body.Bytes(), &todos)
	assert.Len(t, todos, 2, "Expected two todos in the response")
}

func TestGetTodoHandler(t *testing.T) {
	storage.ClearTodos() // Clear the storage before running the test
	storage.AddTodo(&models.Todo{Text: "Test Todo", Completed: false})

	req, err := http.NewRequest("GET", "/todos/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/todos/{id}", GetTodo)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200")
	var todo models.Todo
	json.Unmarshal(rr.Body.Bytes(), &todo)
	assert.Equal(t, "Test Todo", todo.Text, "Expected todo title to match")
	assert.Equal(t, false, todo.Completed, "Expected todo status to match")
}

func TestUpdateTodoHandler(t *testing.T) {
	storage.ClearTodos() // Clear the storage before running the test
	storage.AddTodo(&models.Todo{Text: "Test Todo", Completed: false})
	updatedTodo := models.Todo{Text: "Updated Todo", Completed: false}
	body, _ := json.Marshal(updatedTodo)

	req, err := http.NewRequest("PUT", "/todos/0", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Put("/todos/{id}", UpdateTodo)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200")
	var todo models.Todo
	json.Unmarshal(rr.Body.Bytes(), &todo)
	assert.Equal(t, "Updated Todo", todo.Text, "Expected todo title to match")
	assert.Equal(t, false, todo.Completed, "Expected todo status to match")
}

func TestDeleteTodoHandler(t *testing.T) {
	storage.ClearTodos() // Clear the storage before running the test
	storage.AddTodo(&models.Todo{Text: "Test Todo", Completed: false})

	req, err := http.NewRequest("DELETE", "/todos/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Delete("/todos/{id}", DeleteTodo)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "Expected status code to be 204")
}
