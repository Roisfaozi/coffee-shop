package repository

import "github.com/Roisfaozi/coffee-shop/internal/models"

type UserRepository interface {
	Create(data *models.User) models.User
	Update(data *models.User) models.User
	Delete(data *models.User)
	FindById(userid string) (models.User, error)
	FindAll() []models.User
}
