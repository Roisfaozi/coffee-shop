package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProductHandlerImpl struct {
	productRepo repository.ProductRepository
}

func NewProductHandlerImpl(productRepo repository.ProductRepository) *ProductHandlerImpl {
	return &ProductHandlerImpl{productRepo}
}

func (ph ProductHandlerImpl) CreateProduct(c *gin.Context) {
	var productReq models.ProductRequest
	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productRes, err := ph.productRepo.CreateProduct(c.Request.Context(), &productReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, productRes)
}

func (ph ProductHandlerImpl) UpdateProduct(c *gin.Context) {
	fmt.Println("lele")
	productID := c.Param("id")
	var productReq models.ProductRequest
	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ph.productRepo.UpdateProduct(c.Request.Context(), productID, &productReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (ph ProductHandlerImpl) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	err := ph.productRepo.DeleteProduct(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (ph ProductHandlerImpl) GetAllProducts(c *gin.Context) {
	foodType := c.Query("food_type")

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	products, err := ph.productRepo.GetAllProducts(c.Request.Context(), foodType, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ph ProductHandlerImpl) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	product, err := ph.productRepo.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
