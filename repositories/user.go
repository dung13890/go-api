package repositories

import (
	"go-api/models"
)

type UserRepositoryInterface interface {
	GetAll() models.Users
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

func (repository UserRepository) GetAll() models.Users {
	users := models.Users{
		Users: []models.User{
			models.User{ID: 1, Name: "Oliver1", Email: "Queen@adfa.com1"},
			models.User{ID: 2, Name: "Oliver2", Email: "Queen@adfa.com2"},
			models.User{ID: 3, Name: "Oliver3", Email: "Queen@adfa.com3"},
		},
		Meta: 2,
	}

	return users
}
