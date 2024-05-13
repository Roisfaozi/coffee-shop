package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/test/mocking"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepoMock = mocking.MockUserRepository{}

func expectedMock(method string, returnArg interface{}) {
	userRepoMock.On(method, mock.Anything).Return(returnArg, nil)
}

var reqBody = `{
	"user_id": "123",
	"username": "roisfaozi",
	"password": "roisfaozi",
	"role": "user",
	"email": "rIqFP@example.com"
}`

func TestCreateUser(t *testing.T) {
	r := gin.Default()

	handler := handlers.NewUserHandlerImpl(&userRepoMock)
	r.POST("/user", handler.Create)

	t.Run("Success create user", func(t *testing.T) {

		expectedResult := &config.Result{Message: "Success create user Data"}
		expectedMock("Create", expectedResult)
		req := httptest.NewRequest("POST", "/user", strings.NewReader(reqBody))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, `{"description": "Success create user Data", "status": "Created"}`, w.Body.String())
	})

}

func TestFailedCreateUser(t *testing.T) {
	r := gin.Default()

	handler := handlers.NewUserHandlerImpl(&userRepoMock)
	r.POST("/user", handler.Create)
	var expectedResult *config.Result

	t.Run("Create user empty payload ", func(t *testing.T) {
		expectedResult = &config.Result{Message: "email: non zero value required;password: non zero value required;username: non zero value required"}
		expectedMock("Create", expectedResult)

		var reqBody = `{
							"user_id": "",
							"username": "",
							"password": "",
							"role": "",
							"email": ""
							}`

		req := httptest.NewRequest("POST", "/user", strings.NewReader(reqBody))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.JSONEq(t, `{"description": "email: non zero value required;password: non zero value required;username: non zero value required", "status": "Bad Request"}`, w.Body.String())
	})

	t.Run("Create user empty username", func(t *testing.T) {
		var reqBodyUser = `{
									"user_id": "123",
									"username": "",
									"password": "roisfaozi",
									"role": "user",
									"email": "rIqFP@example.com"
								}`

		expectedResult = &config.Result{Message: "username: non zero value required"}
		expectedMock("Create", expectedResult)

		req := httptest.NewRequest("POST", "/user", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.JSONEq(t, `{"description": "username: non zero value required", "status": "Bad Request"}`, w.Body.String())
	})

	t.Run("Create user without email", func(t *testing.T) {
		var reqBodyUser = `{
									"user_id": "123",
									"username": "roisfaozi",
									"password": "roisfaozi",
									"role": "user",
									"email": ""
								}`

		expectedResult = &config.Result{Message: "email: non zero value required"}
		expectedMock("Create", expectedResult)

		req := httptest.NewRequest("POST", "/user", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.JSONEq(t, `{"description": "email: non zero value required", "status": "Bad Request"}`, w.Body.String())
	})

	t.Run("Create user with wrong email format", func(t *testing.T) {
		var reqBodyUser = `{
									"user_id": "123",
									"username": "roisfaozi",
									"password": "roisfaozi",
									"role": "user",
									"email": "flkdsjfjfs2dfdf.com"
								}`

		expectedResult = &config.Result{Message: "email: flkdsjfjfs2dfdf.com does not validate as email"}
		expectedMock("Create", expectedResult)

		req := httptest.NewRequest("POST", "/user", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.JSONEq(t, `{"description": "email: flkdsjfjfs2dfdf.com does not validate as email", "status": "Bad Request"}`, w.Body.String())
	})

	t.Run("Create user with password less than 6", func(t *testing.T) {
		var reqBodyUser = `{
							"user_id": "123",
							"username": "roisfaozi",
							"password": "rois",
							"role": "user",
							"email": "rIqFP@example.com"
							}`

		expectedResult = &config.Result{Message: "Password minimal 6"}
		expectedMock("Create", expectedResult)

		req := httptest.NewRequest("POST", "/user", strings.NewReader(reqBodyUser))
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.JSONEq(t, `{"description": "Password minimal 6", "status": "Bad Request"}`, w.Body.String())
	})

}

func TestFindUserBy(t *testing.T) {

	r := gin.Default()

	userRepoMock := &userRepoMock

	handler := handlers.NewUserHandlerImpl(userRepoMock)
	r.GET("/user/:id", handler.FindById)
	t.Run("Success find user by id", func(t *testing.T) {
		expectedResult := &config.Result{Data: map[string]interface{}{
			"id":         "2b3be28e-1d40-4dfc-88e3-1db14687aea2",
			"username":   "pakalfjsdl",
			"email":      "makaknan@example.com",
			"created_at": "0001-01-01T00:00:00Z",
			"updated_at": "0001-01-01T00:00:00Z",
		},
			Message: "Success get user Data"}
		expectedMock("FindById", expectedResult)
		req := httptest.NewRequest("GET", "/user/2b3be28e-1d40-4dfc-88e3-1db14687aea2", strings.NewReader("2b3be28e-1d40-4dfc-88e3-1db14687aea2"))
		req.Header.Set("Content-type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.JSONEq(t, `{
        "status": "OK",
        "data": {
            "id": "2b3be28e-1d40-4dfc-88e3-1db14687aea2",
            "username": "pakalfjsdl",
            "email": "makaknan@example.com",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "description": "Success get user Data"
    }`, w.Body.String())
	})
}

func TestFindAllUsers(t *testing.T) {
	r := gin.Default()

	handler := handlers.NewUserHandlerImpl(&userRepoMock)
	r.GET("/user", handler.FindAll)
	expectedResult := &config.Result{
		Data: []map[string]interface{}{
			{
				"id":         "1",
				"username":   "user1",
				"email":      "user1@example.com",
				"created_at": "2024-05-01T00:00:00Z",
				"updated_at": "2024-05-01T00:00:00Z",
			},
			{
				"id":         "2",
				"username":   "user2",
				"email":      "user2@example.com",
				"created_at": "2024-05-02T00:00:00Z",
				"updated_at": "2024-05-02T00:00:00Z",
			},
		},
		Message: "Success get all user Data",
	}
	defer userRepoMock.AssertExpectations(t)

	t.Run("Success", func(t *testing.T) {
		userRepoMock.On("FindAll", mock.Anything).Return(expectedResult, nil)
		req := httptest.NewRequest("GET", "/user", nil)
		req.Header.Set("Content-type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, expectedResult)

		expectedResponse := `{
			"status": "OK",
				   "data": [
				{
					"id": "1",
					"username": "user1",
					"email": "user1@example.com",
					"created_at": "2024-05-01T00:00:00Z",
					"updated_at": "2024-05-01T00:00:00Z"
				},
				{
					"id": "2",
					"username": "user2",
					"email": "user2@example.com",
					"created_at": "2024-05-02T00:00:00Z",
					"updated_at": "2024-05-02T00:00:00Z"
				}
			],
            "description": "Success get all user Data"
        }`

		assert.JSONEq(t, expectedResponse, w.Body.String())

		userRepoMock.AssertCalled(t, "FindAll")
	})

}

func TestFailedFindUserByIde(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := handlers.NewUserHandlerImpl(&userRepoMock)

	t.Run("User not found", func(t *testing.T) {
		mockUserService := &userRepoMock
		mockUserService.On("FindById", mock.Anything).Return(nil, errors.New("User Not Found"))
		k := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/user", handler.FindAll)
		handlers.NewUserHandlerImpl(mockUserService)
		// Create a new HTTP request
		req, err := http.NewRequest("GET", "/user/ere", nil)
		router.ServeHTTP(k, req)
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		router.ServeHTTP(k, req)

		fmt.Println(k.Body.String())

		assert.Equal(t, http.StatusNotFound, k.Code)
		mockUserService.AssertNotCalled(t, "GET", mock.Anything)
	})
}
