package routers

import (
	"github.com/Roisfaozi/coffee-shop/internal/handlers"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/auth")

	repo := repository.NewUserRepositoryImpl(d)
	handler := handlers.NewAuthHandler(repo)

	router.POST("/login", handler.Login)

}
