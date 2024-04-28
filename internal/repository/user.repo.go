package repository

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
)

type UserRepository interface {
	Create(data *models.User) (*config.Result, error)
	FindById(userid string) (*config.Result, error)
	FindAll() (*config.Result, error)
	GetAuthUser(userid string) (*models.User, error)
}
