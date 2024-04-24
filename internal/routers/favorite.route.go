package routers

import (
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/favorite")

	repo := repository.NewFavoriteRepositoryImpl(d)
	handler := handlers.NewFavoriteHandlerImpl(repo)

	router.POST("/", handler.CreateFavorite)
	router.DELETE("/:id", handler.DeleteFavorite)
	router.GET("/:userId", handler.GetFavoritesByUserID)

}
