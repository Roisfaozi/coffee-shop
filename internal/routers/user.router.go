package routers

import (
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/user")

	repo := repository.NewUserRepositoryImpl(d)
	handler := handlers.NewUserHandlerImpl(repo)

	router.POST("/", handler.Create)
	router.PUT("/:id", handler.Update)
	router.DELETE("/:id", handler.Delete)
	router.GET("/", handler.FindAll)
	router.GET("/:id", handler.FindById)
}
