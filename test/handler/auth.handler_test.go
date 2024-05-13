package handler

import (
	"errors"
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"strings"
	"testing"
)

var reqUser = `{
	"username": "Testing-1715461226",
	"password": "roisfaozi",
	}`

func TestLogin(t *testing.T) {
	r := gin.Default()

	handler := handlers.NewAuthHandler(&userRepoMock)
	r.POST("/login", handler.Login)

	t.Run("Success auth user", func(t *testing.T) {
		expectedUser := &models.User{
			ID:       "38961f1d-93ff-459c-ab9b-1e0cd6e0057e",
			Username: "Testing-1715526473",
			Password: "$2a$10$vCQY89Q.WChGT.6fi9VP9eacRJCOYqfkFNPwGU/Gx/5DN0ICKBaCO",
			Role:     "admin",
		}
		expectedMock("GetAuthUser", expectedUser)
		reqBody := `{"username": "Testing-1715526473", "password": "rahasia"}`

		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})

	t.Run("Login without complete payload", func(t *testing.T) {
		var reqBodyUser = `{"username": "","password": "rahasia"}`

		expectedUser := &models.User{
			Username: "Testing-1715526473",
			Password: "rahasia",
		}
		expectedMock("GetAuthUser", expectedUser)

		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)

		assert.JSONEq(t, `{"description": "username: non zero value required", "status": "Bad Request"}`, w.Body.String())
	})

	t.Run("Login with wrong password", func(t *testing.T) {
		var reqBodyUser = `{"username": "Testing-1715526473","password": "rahasias"}`

		expectedUser := &models.User{
			Username: "Testing-1715526473",
			Password: "rahasia",
		}
		expectedMock("GetAuthUser", expectedUser)

		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)

		assert.JSONEq(t, `{"description": "Password is Salah", "status": "Bad Request"}`, w.Body.String())
	})

	t.Run("Login with wrong username", func(t *testing.T) {
		var reqBodyUser = `{"username": "kalemboskueh","password": "rahasia"}`

		userRepoMock.On("GetAuthUser", mock.Anything).Return(nil, errors.New("username not found"))
		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)

		assert.JSONEq(t, `{"description": "username not found", "status": "Unauthorized"}`, w.Body.String())
	})
}
