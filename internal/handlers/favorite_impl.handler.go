package handlers

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/Roisfaozi/coffee-shop/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FavoriteHandlerImpl struct {
	favoriteRepo repository.FavoriteRepository
}

func (fh FavoriteHandlerImpl) CreateFavorite(c *gin.Context) {
	userID := c.Param("userId")

	var request models.FavoriteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
		return
	}
	request.UserID = userID
	favorite, err := fh.favoriteRepo.CreateFavorite(c.Request.Context(), &request)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	pkg.NewRes(http.StatusCreated, favorite).Send(c)
}

func (fh FavoriteHandlerImpl) DeleteFavorite(c *gin.Context) {
	favoriteID := c.Param("id")

	err := fh.favoriteRepo.DeleteFavorite(c.Request.Context(), favoriteID)
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
		Message: "Product updated successfully",
	}).Send(c)
}

func (fh FavoriteHandlerImpl) GetFavoritesByUserID(c *gin.Context) {
	userID := c.Param("userId")

	favorites, err := fh.favoriteRepo.GetFavoritesByUserID(c.Request.Context(), userID)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	pkg.NewRes(http.StatusOK, favorites).Send(c)
}

func NewFavoriteHandlerImpl(favoriteRepo repository.FavoriteRepository) *FavoriteHandlerImpl {
	return &FavoriteHandlerImpl{favoriteRepo}
}
