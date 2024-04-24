package handlers

import "github.com/gin-gonic/gin"

type FavoriteHandler interface {
	CreateFavorite(c *gin.Context)
	DeleteFavorite(c *gin.Context)
	GetFavoritesByUserID(c *gin.Context)
}
