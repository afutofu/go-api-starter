package models

import "github.com/dgrijalva/jwt-go"

/*
User represents a user in the system.

Fields:
  - ID: The unique identifier of the user.
  - Username: The username of the user.
  - Password: The hashed password of the user.
*/
type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
Claims represents the JWT claims for authentication.

Fields:
  - ID: The unique identifier of the authenticated user.
  - Username: The username of the authenticated user.
  - StandardClaims: Embeds the standard JWT claims.
*/
type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
