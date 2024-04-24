package handlers

import (
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteHandlerImpl struct {
	favoriteRepo repository.FavoriteRepository
}

func (fh FavoriteHandlerImpl) CreateFavorite(c *gin.Context) {
	var request models.FavoriteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	favorite, err := fh.favoriteRepo.CreateFavorite(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, favorite)
}

func (fh FavoriteHandlerImpl) DeleteFavorite(c *gin.Context) {
	favoriteID := c.Param("id")

	err := fh.favoriteRepo.DeleteFavorite(c.Request.Context(), favoriteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})
}

func (fh FavoriteHandlerImpl) GetFavoritesByUserID(c *gin.Context) {
	userID := c.Param("userId")

	favorites, err := fh.favoriteRepo.GetFavoritesByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan daftar favorit
	c.JSON(http.StatusOK, favorites)
}

func NewFavoriteHandlerImpl(favoriteRepo repository.FavoriteRepository) *FavoriteHandlerImpl {
	return &FavoriteHandlerImpl{favoriteRepo}
}
