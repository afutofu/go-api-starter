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
	"golang.org/x/crypto/bcrypt"
)

/*
TestRegisterHandler tests the Register handler.
*/
func TestRegisterHandler(t *testing.T) {
	storage.ClearUsers() // Clear the storage before running the test
	user := models.User{Username: "testuser", Password: "password"}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code to be 201")
}

/*
TestLoginHandler tests the Login handler.
*/
func TestLoginHandler(t *testing.T) {
	storage.ClearUsers() // Clear the storage before running the test
	user := models.User{Username: "testuser", Password: "password"}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	storage.SaveUser(&user)

	loginUser := models.User{Username: "testuser", Password: "password"}
	body, _ := json.Marshal(loginUser)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200")
	cookie := rr.Result().Cookies()[0]
	assert.Equal(t, "token", cookie.Name, "Expected cookie name to be 'token'")
}

/*
TestLogoutHandler tests the Logout handler.
*/
func TestLogoutHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Logout)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200")
	cookie := rr.Result().Cookies()[0]
	assert.Equal(t, "token", cookie.Name, "Expected cookie name to be 'token'")
	assert.Equal(t, "", cookie.Value, "Expected cookie value to be empty")
}

/*
TestCreateTodoHandler tests the CreateTodo handler.
*/
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

/*
TestGetTodosHandler tests the GetTodos handler.
*/
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

/*
TestGetTodoHandler tests the GetTodo handler.
*/
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

/*
TestUpdateTodoHandler tests the UpdateTodo handler.
*/
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

/*
TestDeleteTodoHandler tests the DeleteTodo handler.
*/
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
