package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
)

type modelerUser interface {
	InsertUserIntoDB(modelsU.User) error
	DeleteUserFromDB(int) error
	GetAllUsersFromDB() ([]modelsU.User, error)
	FindUserInDB(modelsU.User) (modelsU.User, error)
	GetUser(int) (modelsU.User, error)
}

//UserHandler stradsas
type UserHandler struct {
	M modelerUser
}
