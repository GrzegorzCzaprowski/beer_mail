package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
)

type modelerUser interface {
	InsertUserIntoDB(models.User) error
	DeleteUserFromDB(int) error
	GetAllUsersFromDB() ([]models.User, error)
	FindUserInDB(models.User) (models.User, error)
	GetUser(int) (models.User, error)
}

//UserHandler stradsas
type UserHandler struct {
	M modelerUser
}
