package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
)

type modelerEvent interface {
	InsertEventIntoDB(models.Event) (int, error)
	SendMailsToAllUsers(models.Event, models.User) error
	GetCreator(int) (models.User, error)
	GetAllEventsFromDB() ([]models.Event, error)
	GetEvent(int) (models.Event, error)
	DeleteEventFromDB(int) error
	ConfirmEventForUser(eventID, userID int) error
}

//EventHandler sadasd
type EventHandler struct {
	M modelerEvent
}
