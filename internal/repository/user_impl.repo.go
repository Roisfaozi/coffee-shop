package repository

import (
	"github.com/Roisfaozi/coffee-shop/helper"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	*sqlx.DB
}

func NewUserRepositoryImpl(DB *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (user UserRepositoryImpl) Create(data *models.User) models.User {
	query := `
        INSERT INTO users(username, password, email, role) VALUES(:username,:password, :email, :role) RETURNING id, created_at, updated_at
    `

	_, err := user.NamedExec(query, data)
	helper.PanicIfError(err)
	q := `
        SELECT * FROM users WHERE username=$1
    `

	var foundUser models.User
	err = user.Get(&foundUser, q, data.Username)
	if err != nil {
		return models.User{}
	}
	data.ID = foundUser.ID
	return *data
}

func (user UserRepositoryImpl) Update(data *models.User) models.User {
	query := `
        UPDATE users
SET username = :username,
    email = :email,
    role = :role
WHERE id=:id RETURNING id, username, email, role, created_at, updated_at
    `

	_, err := user.NamedExec(query, data)
	helper.PanicIfError(err)

	return *data
}

func (user UserRepositoryImpl) Delete(data *models.User) {
	query := `
        DELETE FROM users WHERE id=$1
    `

	_, err := user.Exec(query, data.ID)
	helper.PanicIfError(err)
}

func (user UserRepositoryImpl) FindById(userid string) (models.User, error) {
	query := `
        SELECT * FROM users WHERE id=$1
    `

	var foundUser models.User
	err := user.Get(&foundUser, query, userid)
	if err != nil {
		return models.User{}, err
	}
	return foundUser, nil
}

func (user UserRepositoryImpl) FindAll() []models.User {
	query := `
        SELECT * FROM users
    `

	var foundUsers []models.User
	err := user.Select(&foundUsers, query)
	helper.PanicIfError(err)

	return foundUsers
}
