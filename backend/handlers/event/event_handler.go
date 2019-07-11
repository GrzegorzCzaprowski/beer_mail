package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsE"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
)

type modelerEvent interface {
	InsertEventIntoDB(modelsE.Event) (int, error)
	SendMailsToAllUsers(modelsE.Event, modelsU.User) error
	GetCreator(int) (modelsU.User, error)
	GetAllEventsFromDB() ([]modelsE.Event, error)
	GetEvent(int) (modelsE.Event, error)
	DeleteEventFromDB(int) error
	ConfirmEventForUser(int, int, bool) error
}

//EventHandler sadasd
type EventHandler struct {
	M modelerEvent
}
