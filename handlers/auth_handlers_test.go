package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/afutofu/go-api-starter/models"
	"github.com/afutofu/go-api-starter/storage"
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
