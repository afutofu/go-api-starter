package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/afutofu/go-api-starter/models"
	"github.com/afutofu/go-api-starter/storage"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("my_secret_key")

/*
Register handles user registration by storing hashed passwords.

Parameters:
  - w: http.ResponseWriter to write the HTTP response.
  - r: *http.Request containing the registration data.

Returns:
  - void: The function writes the HTTP response directly.
*/
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	storage.SaveUser(&user)

	w.WriteHeader(http.StatusCreated)
}

/*
Login handles user login and issues JWT tokens.

Parameters:
  - w: http.ResponseWriter to write the HTTP response.
  - r: *http.Request containing the login data.

Returns:
  - void: The function writes the HTTP response directly.
*/
func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedUser, ok := storage.GetUserByName(loginUser.Username)
	if !ok || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginUser.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		ID:       storedUser.ID,
		Username: storedUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

/*
Logout handles user logout by expiring the JWT token.

Parameters:
  - w: http.ResponseWriter to write the HTTP response.
  - r: *http.Request

Returns:
  - void: The function writes the HTTP response directly.
*/
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
}
