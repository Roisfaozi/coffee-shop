package pkg

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRes(t *testing.T) {
	// Test case 1: Code < 400, Data and Message present
	data := &config.Result{
		Data:    "test data",
		Message: "test message",
	}
	expected := &Response{
		Code:        200,
		Status:      "OK",
		Data:        "test data",
		Description: "test message",
	}
	assert.Equal(t, expected, NewRes(200, data))

	// Test case 2: Code >= 400, Data present
	data = &config.Result{
		Data: "test error",
	}
	expected = &Response{
		Code:        400,
		Status:      "Bad Request",
		Description: "test error",
	}
	assert.Equal(t, expected, NewRes(400, data))

	// Test case 3: Code >= 400, Message present
	data = &config.Result{
		Message: "test error",
	}
	expected = &Response{
		Code:        400,
		Status:      "Bad Request",
		Description: "test error",
	}
	assert.Equal(t, expected, NewRes(400, data))

	// Test case 4: Code >= 400, No Data or Message present
	data = &config.Result{}
	expected = &Response{
		Code:        400,
		Status:      "Bad Request",
		Description: "Unknown error",
	}
	assert.Equal(t, expected, NewRes(400, data))

	// Test case 5: Meta present
	data = &config.Result{
		Meta: "test meta",
	}
	expected = &Response{
		Code:   200,
		Status: "OK",
		Meta:   "test meta",
	}
	assert.Equal(t, expected, NewRes(200, data))
}

func TestSend(t *testing.T) {
	// Create a new Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Test case 1: Code < 400
	response := &Response{
		Code: 200,
		Data: "test data",
		Meta: "test meta",
	}
	response.Send(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"","data":"test data","meta":"test meta"}`, w.Body.String())
	k := httptest.NewRecorder()
	l, _ := gin.CreateTestContext(k)
	// Test case 2: Code >= 400
	response2 := &Response{
		Code:        400,
		Description: "test error",
		Status:      "Bad Request",
	}
	response2.Send(l)
	assert.Equal(t, http.StatusBadRequest, k.Code)
	assert.Equal(t, "{\"status\":\"Bad Request\",\"description\":\"test error\"}", k.Body.String())
}
