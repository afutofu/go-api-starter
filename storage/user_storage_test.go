package storage

import (
	"testing"

	"github.com/afutofu/go-api-starter/models"
	"github.com/stretchr/testify/assert"
)

/*
TestSaveUserAndGetUserByName tests the SaveUser and GetUserByName functions.
*/
func TestSaveUserAndGetUserByName(t *testing.T) {
	ClearUsers() // Clear the storage before running the test
	user := models.User{Username: "testuser", Password: "password"}
	SaveUser(&user)

	fetchedUser, found := GetUserByName("testuser")
	assert.True(t, found, "Expected user to be found")
	assert.Equal(t, "testuser", fetchedUser.Username, "Expected username to match")
	assert.Equal(t, "password", fetchedUser.Password, "Expected password to match")
}
