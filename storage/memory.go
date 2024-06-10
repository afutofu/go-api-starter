package storage

import (
	"sync"

	"github.com/afutofu/go-api-starter/models"
)

var (
	users      = make(map[int]models.User)
	userByName = make(map[string]int)
	nextID     int
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
	user.ID = nextID
	nextID++
	users[user.ID] = *user
	userByName[user.Username] = user.ID
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
ClearUsers clears all users and resets the ID counter.
*/
func ClearUsers() {
	users = make(map[int]models.User)
	userByName = make(map[string]int)
	nextID = 0
}
