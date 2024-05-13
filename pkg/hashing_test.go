package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	// Test case 1: Successful hash generation
	password := "testpassword"
	hashedPassword, eror := HashPassword(password)
	assert.NoError(t, eror)
	assert.NotEmpty(t, hashedPassword)

	// Test case 2: Error during hash generation
	_, err := HashPassword("")

	assert.Error(t, err)
}

func TestVerifyPassword(t *testing.T) {
	// Test case 1: Successful password verification
	password := "testpassword"
	hashedPassword, _ := HashPassword(password)
	err := VerifyPassword(hashedPassword, password)
	assert.NoError(t, err)

	// Test case 2: Error during password verification
	err = VerifyPassword(hashedPassword, "wrongpassword")
	assert.Error(t, err)

}
