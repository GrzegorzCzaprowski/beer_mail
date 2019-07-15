package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsE"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
)

type modelerEvent interface {
	InsertEvent(modelsE.Event) (int, error)
	SendMails(modelsE.Event, modelsU.User) error
	GetAllEvents() ([]modelsE.Event, error)
	GetEvent(int) (modelsE.Event, error)
	DeleteEvent(int) error
	ConfirmEvent(int, int, bool) error
	GetUser(int) (modelsU.User, error)
	GetUpcomingEvents() ([]modelsE.Event, error)
}

//EventHandler sadasd
type EventHandler struct {
	M modelerEvent
}
