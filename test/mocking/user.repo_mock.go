package mocking

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/stretchr/testify/mock"
)

// Define a mock type for the UserRepositoryImpl struct
type MockUserRepository struct {
	mock.Mock
}

// Implement the Create method of the UserRepositoryImpl interface
func (m *MockUserRepository) Create(data *models.User) (*config.Result, error) {
	args := m.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

// Implement the FindById method of the UserRepositoryImpl interface
func (m *MockUserRepository) FindById(userid string) (*config.Result, error) {
	args := m.Called(userid)
	// first value passed to "Return"
	var r0 *config.Result
	if args.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = args.Get(0).(*config.Result)
	}

	var r1 error

	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}

	return r0, r1
}

// Implement the FindAll method of the UserRepositoryImpl interface
func (m *MockUserRepository) FindAll() (*config.Result, error) {
	args := m.Called()
	return args.Get(0).(*config.Result), args.Error(1)
}

// Implement the GetAuthUser method of the UserRepositoryImpl interface
func (m *MockUserRepository) GetAuthUser(userid string) (*models.User, error) {
	args := m.Called(userid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*models.User), nil
	}
}
