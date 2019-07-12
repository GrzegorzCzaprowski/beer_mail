package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
)

type modelerUser interface {
	InsertUser(modelsU.User) error
	DeleteUser(int) error
	GetAllUsers() ([]modelsU.User, error)
	FindUser(modelsU.User) (modelsU.User, error)
	GetUser(int) (modelsU.User, error)
}

//UserHandler stradsas
type UserHandler struct {
	M modelerUser
}
