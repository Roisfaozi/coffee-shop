package handlers

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/Roisfaozi/coffee-shop/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewUserHandlerImpl(userRepository repository.UserRepository) *UserHandlerImpl {
	return &UserHandlerImpl{userRepository}
}

type UserHandlerImpl struct {
	UserRepository repository.UserRepository
}

func (uh UserHandlerImpl) Create(c *gin.Context) {
	var err error
	user := models.User{
		Role: "user",
	}
	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusUnauthorized, &config.Result{Data: err.Error()}).Send(c)
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusUnauthorized, &config.Result{Data: err.Error()}).Send(c)
		return
	}

	createdUser, err := uh.UserRepository.Create(&user)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusUnauthorized, &config.Result{
			Data:    nil,
			Message: err.Error(),
		}).Send(c)
		return
	}
	pkg.NewRes(http.StatusCreated, createdUser).Send(c)
}

func (uh UserHandlerImpl) FindById(c *gin.Context) {
	userID := c.Param("id")

	user, err := uh.UserRepository.FindById(userID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	pkg.NewRes(200, user).Send(c)
}

func (uh UserHandlerImpl) FindAll(c *gin.Context) {

	users, err := uh.UserRepository.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pkg.NewRes(200, users).Send(c)
}
