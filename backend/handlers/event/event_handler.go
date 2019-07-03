package handlers

type modelerEvent interface {
	CreateEvent()
}

type EventHandler struct {
	M modelerEvent
}
