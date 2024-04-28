package routers

import (
	"github.com/Roisfaozi/coffee-shop/exception"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(exception.GlobalErrorHandler())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	user(router, db)
	products(router, db)
	favorite(router, db)
	auth(router, db)
	return router

}
