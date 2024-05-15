package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/test/mocking"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProductById(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	var productRepoMock = mocking.MockProductRepository{}
	t.Run("Success get product by id", func(t *testing.T) {
		// Mock Product
		producId := "6ba10979-8f56-4e01-aaf9-56a492e51898"
		expectedResult := `{
		"status": "OK",
			"data": {
				"id": "6ba10979-8f56-4e01-aaf9-56a492e51898",
				"name": "Nasi Goreng",
				"price": 30000,
				"currency": "IDR",
				"description": "Deskripsi produk kopi yang luar biasa",
				"image_url": "",
				"category": "Makanan",
				"created_at": "2024-05-14T02:13:44.952317Z",
				"updated_at": "2024-05-14T02:13:44.952317Z",
				"sizes": [
					{
						"id": "ac53f05a-2d88-4562-bef1-301d32dff13e",
						"size_name": "R"
					},
					{
						"id": "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
						"size_name": "L"
					},
					{
						"id": "6ab889b0-a10d-4f92-ab26-d70710205a9e",
						"size_name": "XL"
					}
				]
			}
		}`
		mockProduct := &config.Result{
			Data: map[string]interface{}{
				"id":          "6ba10979-8f56-4e01-aaf9-56a492e51898",
				"name":        "Nasi Goreng",
				"price":       30000,
				"currency":    "IDR",
				"description": "Deskripsi produk kopi yang luar biasa",
				"image_url":   "",
				"category":    "Makanan",
				"created_at":  "2024-05-14T02:13:44.952317Z",
				"updated_at":  "2024-05-14T02:13:44.952317Z",
				"sizes": []map[string]interface{}{
					{
						"id":        "ac53f05a-2d88-4562-bef1-301d32dff13e",
						"size_name": "R",
					},
					{
						"id":        "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
						"size_name": "L",
					},
					{
						"id":        "6ab889b0-a10d-4f92-ab26-d70710205a9e",
						"size_name": "XL",
					},
				},
			},
		}

		mockProductRepo := &productRepoMock
		mockProductRepo.On("GetProductByID", mock.Anything, producId).Return(mockProduct, nil)

		// Setup Router
		router := gin.Default()
		handler := handlers.NewProductHandlerImpl(mockProductRepo)
		router.GET("/products/:id", handler.GetProductByID)
		// Make Request
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", producId), nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Assert Response
		assert.Equal(t, http.StatusOK, rr.Code)
		var responseBody map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.NotNil(t, responseBody["data"])
		assert.Equal(t, producId, responseBody["data"].(map[string]interface{})["id"])
		assert.JSONEq(t, expectedResult, rr.Body.String())
	})

	t.Run("NotFound", func(t *testing.T) {
		// Mock Product ID
		productId := "dfdfdfid-salahe453465"
		expectedResult := `{
			"status": "Not Found",
			"description": "Product not found"
		}`

		mockProductRepo := &productRepoMock
		mockProductRepo.On("GetProductByID", mock.Anything, productId).Return(nil, errors.New("Product not found"))
		// Setup Router
		router := gin.Default()
		handler := handlers.NewProductHandlerImpl(mockProductRepo)
		router.GET("/products/:id", handler.GetProductByID)

		// Make Request
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", productId), nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Assert Response
		assert.Equal(t, http.StatusNotFound, rr.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "Not Found", responseBody["status"])
		assert.Equal(t, "Product not found", responseBody["description"])
		assert.JSONEq(t, expectedResult, rr.Body.String())
	})

}

func TestFailedUpdateProduct(t *testing.T) {
	// Membuat mock untuk repository produk
	var productRepoMock = mocking.MockProductRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	// Konten body permintaan
	var reqBody = `{
        "name": "Updated Product",
        "description": "This is an updated product",
        "price": 35000,
        "currency": "USD",
        "category": "Minuman"
    }`

	productID := "6ba10979-8f56-4e01-aaf9-56a492e51898"
	// Pengujian kasus "Failed update product when product not found"
	t.Run("Failed update product when product not found", func(t *testing.T) {
		// Hasil yang diharapkan dari pemanggilan UpdateProduct
		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(nil, errors.New("Product not found"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler UpdateProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}
			ctx.Set("image_url", "https://example.com/updated-product.png")

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.UpdateProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, server.URL+"/products/"+productID, strings.NewReader(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Equal(t, "Not Found", responseBody["status"])
		assert.Equal(t, "Product not found", responseBody["description"])
	})

	// Pengujian kasus "Failed update product due to repository error"
	t.Run("Failed update product due to repository error", func(t *testing.T) {

		// Hasil yang diharapkan dari pemanggilan UpdateProduct
		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(nil, nil)
		productRepoMock.On("UpdateProduct", mock.Anything, productID, mock.Anything).Return(errors.New("repository error"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler UpdateProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}
			ctx.Set("image_url", "https://example.com/updated-product.png")

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.UpdateProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, server.URL+"/products/"+productID, strings.NewReader(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", http.StatusInternalServerError, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, "Internal Server Error", responseBody["status"])
	})

}

func TestCreateProduct(t *testing.T) {
	// Membuat mock untuk repository produk
	var productRepoMock = mocking.MockProductRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	// Konten body permintaan
	var reqBody = `{
        "name": "Test Product",
        "description": "This is a test product",
        "price": 30000,
        "currency": "IDR",
        "category": "Makanan"
    }`

	// Pengujian kasus "Success create product"
	t.Run("Success create product", func(t *testing.T) {
		// Hasil yang diharapkan dari pemanggilan CreateProduct
		expectedResult := &config.Result{
			Message: "Product created successfully",
			Data: map[string]interface{}{
				"id": "6ba10979-8f56-4e01-aaf9-56a492e51898",
			},
		}

		// Mendaftarkan harapan untuk pemanggilan CreateProduct
		productRepoMock.On("CreateProduct", mock.Anything, mock.Anything).Return(expectedResult, nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler CreateProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Set("image_url", "https://example.com/test-product.png")

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.CreateProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		resp, err := http.Post(server.URL+"/products", "application/json", strings.NewReader(reqBody))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		assert.Equal(t, "Created", responseBody["status"])

		assert.Equal(t, expectedResult.Message, responseBody["description"])
		assert.Equal(t, expectedResult.Data, responseBody["data"])

	})

	t.Run("failed create product with empty required fields", func(t *testing.T) {
		// Hasil yang diharapkan dari pemanggilan CreateProduct
		expectedResult := &config.Result{
			Message: "Key: 'ProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'ProductRequest.Price' Error:Field validation for 'Price' failed on the 'required' tag\nKey: 'ProductRequest.Currency' Error:Field validation for 'Currency' failed on the 'required' tag",
		}

		var reqBody = `{
        "name": "",
        "description": "",
        "currency": "",
        "category": "Makanan"
    }`

		// Mendaftarkan harapan untuk pemanggilan CreateProduct
		productRepoMock.On("CreateProduct", mock.Anything, mock.Anything).Return(expectedResult, nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler CreateProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Set("image_url", "https://example.com/test-product.png")

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.CreateProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		resp, err := http.Post(server.URL+"/products", "application/json", strings.NewReader(reqBody))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		assert.Equal(t, "Bad Request", responseBody["status"])
		assert.Equal(t, expectedResult.Message, responseBody["description"])

	})

}

func TestUpdateProduct(t *testing.T) {
	// Membuat mock untuk repository produk
	var productRepoMock = mocking.MockProductRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	// Konten body permintaan
	var reqBody = `{
        "name": "Updated Product",
        "description": "This is an updated product",
        "price": 35000,
        "currency": "USD",
        "category": "Minuman"
    }`

	productID := "6ba10979-8f56-4e01-aaf9-56a492e51898"

	// Pengujian kasus "Success update product"
	t.Run("Success update product", func(t *testing.T) {

		mockProduct := &config.Result{
			Data: map[string]interface{}{
				"id":          "6ba10979-8f56-4e01-aaf9-56a492e51898",
				"name":        "Nasi Goreng",
				"price":       30000,
				"currency":    "IDR",
				"description": "Deskripsi produk kopi yang luar biasa",
				"image_url":   "",
				"category":    "Makanan",
				"created_at":  "2024-05-14T02:13:44.952317Z",
				"updated_at":  "2024-05-14T02:13:44.952317Z",
				"sizes": []map[string]interface{}{
					{
						"id":        "ac53f05a-2d88-4562-bef1-301d32dff13e",
						"size_name": "R",
					},
					{
						"id":        "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
						"size_name": "L",
					},
					{
						"id":        "6ab889b0-a10d-4f92-ab26-d70710205a9e",
						"size_name": "XL",
					},
				},
			},
		}

		// Hasil yang diharapkan dari pemanggilan UpdateProduct
		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(mockProduct, nil)
		productRepoMock.On("UpdateProduct", mock.Anything, productID, mock.Anything).Return(nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler UpdateProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}
			ctx.Set("image_url", "https://example.com/updated-product.png")

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.UpdateProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, server.URL+"/products/"+productID, strings.NewReader(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "OK", responseBody["status"])
		assert.Equal(t, "Product updated successfully", responseBody["description"])
	})

	// Pengujian kasus "Failed update product with invalid fields"
	t.Run("Failed update product with invalid fields", func(t *testing.T) {
		var invalidReqBody = `{
            "name": "",
            "description": "",
            "price": 0,
            "currency": "",
            "category": "Minuman"
        }`

		// Mendaftarkan harapan untuk pemanggilan GetProductByID
		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(&models.ProductRequest{}, nil)
		productRepoMock.On("UpdateProduct", mock.Anything, productID, mock.Anything).Return(nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler UpdateProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}
			ctx.Set("image_url", "https://example.com/updated-product.png")

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.UpdateProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, server.URL+"/products/"+productID, strings.NewReader(invalidReqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, "Bad Request", responseBody["status"])
		assert.Contains(t, responseBody["description"], "Field validation for 'Name' failed")
	})

}

func TestDeleteProduct(t *testing.T) {
	// Membuat mock untuk repository produk
	var productRepoMock = mocking.MockProductRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	productID := "6ba10979-8f56-4e01-aaf9-56a492e51898"

	// Pengujian kasus "Success delete product"
	t.Run("Success delete product", func(t *testing.T) {
		mockProduct := &config.Result{
			Data: map[string]interface{}{
				"id":          "6ba10979-8f56-4e01-aaf9-56a492e51898",
				"name":        "Nasi Goreng",
				"price":       30000,
				"currency":    "IDR",
				"description": "Deskripsi produk kopi yang luar biasa",
				"image_url":   "",
				"category":    "Makanan",
				"created_at":  "2024-05-14T02:13:44.952317Z",
				"updated_at":  "2024-05-14T02:13:44.952317Z",
				"sizes": []map[string]interface{}{
					{
						"id":        "ac53f05a-2d88-4562-bef1-301d32dff13e",
						"size_name": "R",
					},
					{
						"id":        "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
						"size_name": "L",
					},
					{
						"id":        "6ab889b0-a10d-4f92-ab26-d70710205a9e",
						"size_name": "XL",
					},
				},
			},
		}

		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(mockProduct, nil)
		productRepoMock.On("DeleteProduct", mock.Anything, productID).Return(nil)

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler DeleteProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.DeleteProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/products/"+productID, nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "OK", responseBody["status"])
		assert.Equal(t, "Product deleted successfully", responseBody["description"])
	})

	// Pengujian kasus "Failed to delete product"
	t.Run("Failed to delete product", func(t *testing.T) {

		mockProduct := &config.Result{
			Data: map[string]interface{}{
				"id":          "6ba10979-8f56-4e01-aaf9-56a492e51898",
				"name":        "Nasi Goreng",
				"price":       30000,
				"currency":    "IDR",
				"description": "Deskripsi produk kopi yang luar biasa",
				"image_url":   "",
				"category":    "Makanan",
				"created_at":  "2024-05-14T02:13:44.952317Z",
				"updated_at":  "2024-05-14T02:13:44.952317Z",
				"sizes": []map[string]interface{}{
					{
						"id":        "ac53f05a-2d88-4562-bef1-301d32dff13e",
						"size_name": "R",
					},
					{
						"id":        "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
						"size_name": "L",
					},
					{
						"id":        "6ab889b0-a10d-4f92-ab26-d70710205a9e",
						"size_name": "XL",
					},
				},
			},
		}

		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(mockProduct, nil)
		productRepoMock.On("DeleteProduct", mock.Anything, productID).Return(errors.New("failed to delete product"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler DeleteProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.DeleteProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/products/"+productID, nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", http.StatusInternalServerError, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "Internal Server Error", responseBody["status"])
		assert.Equal(t, "failed to delete product", responseBody["description"])
	})

	// Pengujian kasus "Product not found"
	t.Run("Product not found", func(t *testing.T) {
		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(nil, errors.New("Product not found"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler DeleteProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.DeleteProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/products/"+productID, nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "Not Found", responseBody["status"])
		assert.Equal(t, "Product not found", responseBody["description"])
	})

}

func TestDeleteProductNotFound(t *testing.T) {
	// Membuat mock untuk repository produk
	var productRepoMock = mocking.MockProductRepository{}

	// Mode test untuk Gin
	gin.SetMode(gin.TestMode)

	productID := "6ba10979-8f56-4e01-aaf9-56a492e51898"

	// Pengujian kasus "Product not found"
	t.Run("Product not found", func(t *testing.T) {
		productRepoMock.On("GetProductByID", mock.Anything, productID).Return(nil, errors.New("Product not found"))

		// Membuat server HTTP sementara
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Menjalankan handler DeleteProduct dengan request yang diberikan
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = gin.Params{gin.Param{Key: "id", Value: productID}}

			handler := handlers.NewProductHandlerImpl(&productRepoMock)
			handler.DeleteProduct(ctx)
		}))
		defer server.Close()

		// Membuat permintaan HTTP ke server
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/products/"+productID, nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Memeriksa respons
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}

		// Membaca dan memeriksa body respons
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "Not Found", responseBody["status"])
		assert.Equal(t, "Product not found", responseBody["description"])
	})

}
