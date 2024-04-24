package routers

import (
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/products")

	repo := repository.NewProductRepositoryImpl(d)
	handler := handlers.NewProductHandlerImpl(repo)

	router.POST("/", handler.CreateProduct)
	router.PUT("/:id", handler.UpdateProduct)
	router.DELETE("/:id", handler.DeleteProduct)
	router.GET("/", handler.GetAllProducts)
	router.GET("/:id", handler.GetProductByID)

}
