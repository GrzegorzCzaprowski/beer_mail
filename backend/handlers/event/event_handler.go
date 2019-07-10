package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
)

type modelerEvent interface {
	InsertEventIntoDB(models.Event) error
	SendMailsToAllUsers(models.Event, models.User) error
	GetCreator(int) (models.User, error)
	GetAllEventsFromDB() ([]models.Event, error)
}

//EventHandler sadasd
type EventHandler struct {
	M modelerEvent
}
