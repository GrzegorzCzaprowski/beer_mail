package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
)

type modelerUser interface {
	CreateUser()
	DeleteUser()
	GetAllUsers()
	FindUserInDB(models.User) (models.User, error)
	Logout()
}

type UserHandler struct {
	M modelerUser
}
