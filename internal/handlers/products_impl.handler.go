package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/pkg"

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
	if err := c.ShouldBind(&productReq); err != nil {
		log.Println("CreateProduct", err)
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	productReq.ImageURL = c.MustGet("image_url").(string)
	productRes, err := ph.productRepo.CreateProduct(c.Request.Context(), &productReq)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	pkg.NewRes(http.StatusCreated, productRes).Send(c)
}

func (ph ProductHandlerImpl) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var productReq models.ProductRequest
	if err := c.ShouldBind(&productReq); err != nil {
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	_, err := ph.productRepo.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusNotFound, &config.Result{
			Data:    nil,
			Message: "Product not found",
		}).Send(c)
		return
	}

	productReq.ImageURL = c.MustGet("image_url").(string)

	err = ph.productRepo.UpdateProduct(c.Request.Context(), productID, &productReq)
	if err != nil {
		log.Println(productReq)
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	pkg.NewRes(http.StatusOK, &config.Result{
		Data:    nil,
		Message: " Product updated successfully",
	}).Send(c)
}

func (ph ProductHandlerImpl) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	_, err := ph.productRepo.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusNotFound, &config.Result{
			Data:    nil,
			Message: "Product not found",
		}).Send(c)
		return
	}

	err = ph.productRepo.DeleteProduct(c.Request.Context(), productID)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	pkg.NewRes(http.StatusOK, &config.Result{
		Data:    nil,
		Message: "Product deleted successfully",
	}).Send(c)

}

func (ph ProductHandlerImpl) GetAllProducts(c *gin.Context) {
	foodType := c.Query("food_type")

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		log.Println(err)
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)

		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		log.Println(err)
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}

	products, err := ph.productRepo.GetAllProducts(c.Request.Context(), foodType, page, limit)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}

	pkg.NewRes(http.StatusOK, products).Send(c)
}

func (ph ProductHandlerImpl) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	product, err := ph.productRepo.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusNotFound, &config.Result{
			Data:    nil,
			Message: "Product not found",
		}).Send(c)
		return
	}

	pkg.NewRes(http.StatusOK, product).Send(c)
}
