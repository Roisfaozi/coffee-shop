package handler

import (
	"encoding/json"
	"errors"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/test/mocking"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateFavorite(t *testing.T) {
	// Membuat mock untuk repository favorite
	var favoriteRepoMock = mocking.MockFavoriteRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	// Konten body permintaan
	var reqBody = `{
        "user_id": "user123",
        "product_id": "6f70260b-b7b2-4fdd-a65b-a88782d5dd4b"
    }`

	// Pengujian kasus "Success create favorite"
	t.Run("Success create favorite", func(t *testing.T) {
		// Hasil yang diharapkan dari pemanggilan CreateFavorite
		expectedResult := &config.Result{
			Message: "Favorite created successfully",
			Data: map[string]interface{}{
				"id": "66cfea41-2e16-421e-91c2-d6741c83a4b5",
			},
		}

		// Mendaftarkan harapan untuk pemanggilan CreateFavorite
		favoriteRepoMock.On("CreateFavorite", mock.Anything, mock.Anything).Return(expectedResult, nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler CreateFavorite dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
			handler.CreateFavorite(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		resp, err := http.Post(server.URL+"/favorites", "application/json", strings.NewReader(reqBody))
		require.NoError(t, err)
		defer resp.Body.Close()

		// Memeriksa respons
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		// Membaca dan memeriksa body respons
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "Created", responseBody["status"])
		assert.Equal(t, expectedResult.Message, responseBody["description"])
		assert.Equal(t, expectedResult.Data, responseBody["data"])
	})

}

func TestFailedCreateFavorite(t *testing.T) {
	// Membuat mock untuk repository favorite
	var favoriteRepoMock = mocking.MockFavoriteRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	t.Run("failed create favorite with empty required fields", func(t *testing.T) {
		// Hasil yang diharapkan dari pemanggilan CreateFavorite
		expectedResult := errors.New("Key: 'FavoriteRequest.ProductID' Error:Field validation for 'ProductID' failed on the 'required' tag")

		var reqBody = `{
            "product_id": ""
        }`

		// Mendaftarkan harapan untuk pemanggilan CreateFavorite
		favoriteRepoMock.On("CreateFavorite", mock.Anything, mock.Anything).Return(nil, expectedResult)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler CreateFavorite dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
			handler.CreateFavorite(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		resp, err := http.Post(server.URL+"/favorites", "application/json", strings.NewReader(reqBody))
		require.NoError(t, err)
		defer resp.Body.Close()

		// Memeriksa respons
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		// Membaca dan memeriksa body respons
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "Bad Request", responseBody["status"])
		assert.Equal(t, expectedResult.Error(), responseBody["description"])
	})
}

func TestDeleteFavorite(t *testing.T) {
	// Membuat mock untuk repository favorite
	var favoriteRepoMock = mocking.MockFavoriteRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	t.Run("success delete favorite", func(t *testing.T) {
		favoriteID := "6ba10979-8f56-4e01-aaf9-56a492e51898"

		// Mendaftarkan harapan untuk pemanggilan GetFavoritesByUserID dan DeleteFavorite
		favoriteRepoMock.On("DeleteFavorite", mock.Anything, favoriteID).Return(nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler DeleteFavorite dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: favoriteID}}

			handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
			handler.DeleteFavorite(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/favorites/"+favoriteID, nil)
		require.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		// Memeriksa respons
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Membaca dan memeriksa body respons
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "OK", responseBody["status"])
		assert.Equal(t, "Favorite deleted successfully", responseBody["description"])
	})

	t.Run("failed delete favorite with internal error", func(t *testing.T) {
		favoriteID := "6ba10979-8f56-4e01-aaf9-56a492e51898"

		// Mendaftarkan harapan untuk pemanggilan GetFavoritesByUserID dan DeleteFavorite
		favoriteRepoMock.On("DeleteFavorite", mock.Anything, favoriteID).Return(errors.New("internal server error"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler DeleteFavorite dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: favoriteID}}

			handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
			handler.DeleteFavorite(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/favorites/"+favoriteID, nil)
		require.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		// Memeriksa respons
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		// Membaca dan memeriksa body respons
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "Internal Server Error", responseBody["status"])
		assert.Equal(t, "internal server error", responseBody["description"])
	})

	//
	//t.Run("failed delete favorite with non-existent ID", func(t *testing.T) {
	//	favoriteID := "non-existent-id"
	//
	//	// Mendaftarkan harapan untuk pemanggilan GetFavoritesByUserID
	//	favoriteRepoMock.On("DeleteFavorite", mock.Anything, favoriteID).Return(nil, errors.New("Favorite not found"))
	//
	//	// Membuat server HTTP sementara
	//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	//		// Menjalankan handler DeleteFavorite dengan request yang diberikan
	//		ctx, _ := gin.CreateTestContext(w)
	//		ctx.Request = req
	//		ctx.Params = gin.Params{gin.Param{Key: "id", Value: favoriteID}}
	//
	//		handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
	//		handler.DeleteFavorite(ctx)
	//	}))
	//	defer server.Close()
	//
	//	// Membuat permintaan HTTP ke server
	//	req, err := http.NewRequest(http.MethodDelete, server.URL+"/favorites/"+favoriteID, nil)
	//	require.NoError(t, err)
	//
	//	resp, err := http.DefaultClient.Do(req)
	//	require.NoError(t, err)
	//	defer resp.Body.Close()
	//
	//	// Memeriksa respons
	//	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	//
	//	// Membaca dan memeriksa body respons
	//	body, err := ioutil.ReadAll(resp.Body)
	//	require.NoError(t, err)
	//
	//	var responseBody map[string]interface{}
	//	err = json.Unmarshal(body, &responseBody)
	//	require.NoError(t, err)
	//
	//	assert.Equal(t, "Not Found", responseBody["status"])
	//	assert.Equal(t, "Favorite not found", responseBody["description"])
	//})

}

func TestGetFavoritesByUserID(t *testing.T) {
	// Membuat mock untuk repository favorite
	var favoriteRepoMock = mocking.MockFavoriteRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	t.Run("success get favorites by user ID", func(t *testing.T) {
		userID := "6ba10979-8f56-4e01-aaf9-56a492e51898"

		// Mock data
		expectedFavorites := &config.Result{
			Data: []models.Favorite{
				{ID: "1", UserID: userID, ProductID: "101"},
				{ID: "2", UserID: userID, ProductID: "102"},
			},
			Message: "Favorites retrieved successfully",
		}

		// Mendaftarkan harapan untuk pemanggilan GetFavoritesByUserID
		favoriteRepoMock.On("GetFavoritesByUserID", mock.Anything, mock.Anything).Return(expectedFavorites, nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler GetFavoritesByUserID dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: userID}}

			handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
			handler.GetFavoritesByUserID(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		resp, err := http.Get(server.URL + "/favorites/" + userID)
		require.NoError(t, err)
		defer resp.Body.Close()

		// Memeriksa respons
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Membaca dan memeriksa body respons
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "OK", responseBody["status"])
		assert.Equal(t, "Favorites retrieved successfully", responseBody["description"])
		assert.NotNil(t, responseBody["data"])
	})
	t.Run("failed get favorites by user ID with non-existent user", func(t *testing.T) {
		userID := "non-existent-id"

		// Mendaftarkan harapan untuk pemanggilan GetFavoritesByUserID
		favoriteRepoMock.On("GetFavoritesByUserID", mock.Anything, mock.Anything).Return(nil, errors.New("user not found"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler GetFavoritesByUserID dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: userID}}

			handler := handlers.NewFavoriteHandlerImpl(&favoriteRepoMock)
			handler.GetFavoritesByUserID(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		resp, err := http.Get(server.URL + "/favorites/" + userID)
		require.NoError(t, err)
		defer resp.Body.Close()

		// Memeriksa respons
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		// Membaca dan memeriksa body respons
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "Internal Server Error", responseBody["status"])
		assert.Equal(t, "user not found", responseBody["description"])
	})
}
