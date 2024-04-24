package handlers

import (
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserHandlerImpl(userRepository repository.UserRepository) *UserHandlerImpl {
	return &UserHandlerImpl{userRepository}
}

type UserHandlerImpl struct {
	UserRepository repository.UserRepository
}

func (uh UserHandlerImpl) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser := uh.UserRepository.Create(&user)
	userResponse := &models.UserResponse{
		createdUser.ID,
		createdUser.Username,
		createdUser.Email,
		createdUser.Role,
	}
	c.JSON(http.StatusCreated, userResponse)
}

func (uh UserHandlerImpl) Update(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := uh.UserRepository.FindById(userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID
	user = uh.UserRepository.Update(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uh UserHandlerImpl) Delete(c *gin.Context) {
	userID := c.Param("id")

	user, err := uh.UserRepository.FindById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	uh.UserRepository.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uh UserHandlerImpl) FindById(c *gin.Context) {
	userID := c.Param("id")

	user, err := uh.UserRepository.FindById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userResponse := &models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	c.JSON(http.StatusOK, userResponse)
}

func (uh UserHandlerImpl) FindAll(c *gin.Context) {

	users := uh.UserRepository.FindAll()
	var userResponses []*models.UserResponse
	for _, user := range users {
		userResponse := &models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}
		userResponses = append(userResponses, userResponse)
	}
	c.JSON(http.StatusOK, userResponses)
}
