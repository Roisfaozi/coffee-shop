package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewToken(t *testing.T) {
	// Test NewToken function with valid input
	token := NewToken("123", "admin")
	assert.NotNil(t, token)
	assert.Equal(t, "123", token.Id)
	assert.Equal(t, "admin", token.Role)

	// Test NewToken function with empty ID
	token = NewToken("", "admin")
	assert.Equal(t, "", token.Id)
	assert.Equal(t, "admin", token.Role)

	// Test NewToken function with empty role
	token = NewToken("123", "")

	assert.Equal(t, "123", token.Id)
	assert.Equal(t, "", token.Role)
}

func TestGenerateToken(t *testing.T) {
	// Mock the environment variable for testing
	os.Setenv("JWT_KEYS", "test_secret_key")

	// Test Generate function with a valid claim
	claim := &claims{Id: "123", Role: "admin"}
	tokenString, err := claim.Generate()
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)

}

func TestGenerateTokenEmpty(t *testing.T) {
	// Mock the environment variable for testing
	os.Setenv("JWT_KEYS", "test_secret_key")

	// Test Generate function with an empty claim
	claim := &claims{}
	tokenString, err := claim.Generate()
	assert.Error(t, err)
	assert.Empty(t, tokenString)
}

func TestVerifyToken(t *testing.T) {
	// Mock the environment variable for testing
	os.Setenv("JWT_KEYS", "test_secret_key")

	claim := &claims{Id: "123", Role: "admin"}
	tokenString, err := claim.Generate()

	// Test VerifyToken function with a valid token
	validToken := tokenString
	claimData, err := VerifyToken(validToken)
	assert.NoError(t, err)
	assert.NotNil(t, claimData)

	// Test VerifyToken function with an invalid token
	invalidToken := "invalid_token_string"
	claim, err = VerifyToken(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, claim)

	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImYwMTMzMjg1LWQ2MzAtNDgzYS1iYzM0LWMzZjJjYjAwMTFkYiIsInJvbGUiOiJhZG1pbiIsImlzcyI6ImJhY2tHb2xhbmciLCJleHAiOjE3MTUzNDk5ODd9.tad7MguC1pecQGZHsK-G4zK8nB23wms2gUQZenKH_Fs"
	// Test VerifyToken function with an expired token
	claim, err = VerifyToken(expiredToken)
	assert.Error(t, err)
	assert.Nil(t, claim)
}
