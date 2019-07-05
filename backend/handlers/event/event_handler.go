package handlers

import (
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
)

type modelerEvent interface {
	InsertEventIntoDB(models.Event) error
}

type EventHandler struct {
	M modelerEvent
}
