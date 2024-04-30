package routers

import (
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/middleware"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/products")

	repo := repository.NewProductRepositoryImpl(d)
	handler := handlers.NewProductHandlerImpl(repo)

	router.POST("/", middleware.Authjwt("admin"), middleware.Upload, handler.CreateProduct)
	router.PUT("/:id", middleware.Authjwt("admin"), middleware.Upload, handler.UpdateProduct)
	router.DELETE("/:id", middleware.Authjwt("admin"), handler.DeleteProduct)
	router.GET("/", handler.GetAllProducts)
	router.GET("/:id", middleware.Authjwt("admin", "user"), handler.GetProductByID)

}
