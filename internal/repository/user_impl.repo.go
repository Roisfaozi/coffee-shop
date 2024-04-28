package repository

import (
	"errors"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/helper"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserRepositoryImpl struct {
	*sqlx.DB
}

func NewUserRepositoryImpl(DB *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (user UserRepositoryImpl) Create(data *models.User) (*config.Result, error) {
	query := `
        INSERT INTO users(username, password, email, role) VALUES(:username,:password, :email, :role) RETURNING id, created_at, updated_at
    `

	_, err := user.NamedExec(query, data)
	helper.PanicIfError(err)
	q := `
        SELECT id, username, email FROM users WHERE username=$1
    `

	var foundUser models.User
	err = user.Get(&foundUser, q, data.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &config.Result{Data: models.UserResponse{ID: foundUser.ID,
		Username: foundUser.Username,
		Email:    foundUser.Email,
	}, Message: "1 data user created"}, nil
}

func (user UserRepositoryImpl) FindById(userid string) (*config.Result, error) {
	query := `
        SELECT id, username, email FROM users WHERE id=$1
    `

	var foundUser models.User
	err := user.Get(&foundUser, query, userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &config.Result{Data: foundUser, Message: "Success get user Data"}, nil
}

func (user UserRepositoryImpl) FindAll() (*config.Result, error) {
	query := `
         SELECT id, username, email FROM users
    `
	var foundUsers []models.User
	err := user.Select(&foundUsers, query)
	if err != nil {
		return nil, err
	}
	return &config.Result{Data: foundUsers, Message: "Success get all user Data"}, nil
}

func (user UserRepositoryImpl) GetAuthUser(userid string) (*models.User, error) {
	var result models.User
	q := `SELECT id, username, role, password FROM users WHERE username =?`

	if err := user.Get(&result, user.Rebind(q), userid); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}
		return nil, err
	}

	return &result, nil
}
