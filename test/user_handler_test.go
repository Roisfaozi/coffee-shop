package test

import (
	"bytes"
	"fmt"
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupDB() (*sqlx.DB, error) {
	//host := os.Getenv("127.0.0.1")
	//user := os.Getenv("rois")
	//password := os.Getenv("rois")
	//dbname := os.Getenv("go-coffee-shop-test")
	//fmt.Println(user)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", "rois", "rois", "go-coffee-shop-test")

	return sqlx.Connect("postgres", config)
}

func setupRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	userRepo := repository.NewUserRepositoryImpl(db)
	userHandler := handlers.NewUserHandlerImpl(userRepo)
	router.POST("/user", userHandler.Create)
	return router
}

func TruncateCategory(db *sqlx.DB) {
	_, err := db.Exec("truncate users")
	if err != nil {
		return
	}
}

func TestCreateUserSuccess(t *testing.T) {
	// Setup database pengujian
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Setup router
	router := setupRouter(db)

	// Bersihkan data pengguna sebelum pengujian
	TruncateCategory(db)

	// Buat payload untuk request pengguna yang berhasil
	payload := []byte(`{"username": "testuser", "email": "test@example.com", "role": "user"}`)

	// Buat request HTTP dengan payload
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Eksekusi request dengan router
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Periksa kode status yang diharapkan (201 Created)
	assert.Equal(t, http.StatusCreated, resp.Code)

	// Periksa apakah pengguna telah dibuat dalam database
	// Anda dapat menyesuaikan implementasi ini sesuai dengan logika aplikasi Anda
	// Di sini, kita hanya memastikan bahwa tidak ada kesalahan yang terjadi
	// dan tidak ada pesan kesalahan yang dikembalikan
	assert.NotEmpty(t, resp.Body.String())
}

func TestCreateUserFailed(t *testing.T) {
	// Setup database pengujian
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Setup router
	router := setupRouter(db)

	// Bersihkan data pengguna sebelum pengujian
	TruncateCategory(db)

	// Buat payload untuk request pengguna yang gagal (tanpa email)
	payload := []byte(`{"username": "testuser", "role": "user"}`)

	// Buat request HTTP dengan payload
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Eksekusi request dengan router
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Periksa kode status yang diharapkan (400 Bad Request)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Periksa apakah pesan kesalahan yang diharapkan dikembalikan dalam body respons
	expected := `{"error":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`
	assert.Equal(t, expected, resp.Body.String())
}
