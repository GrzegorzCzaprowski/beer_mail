package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
)

type modelerEvent interface {
	InsertEventIntoDB(models.Event) error
	SendMailsToAllUsers(models.Event, models.User) error
	GetCreator(int) (models.User, error)
	GetAllEventsFromDB() ([]models.Event, error)
	GetEvent(int) (models.Event, error)
	DeleteEventFromDB(int) error
}

//EventHandler sadasd
type EventHandler struct {
	M modelerEvent
}
