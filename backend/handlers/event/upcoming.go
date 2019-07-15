package handlers

import (
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/errorHandler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h EventHandler) Upcoming(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := authorization.UserAuthentication(w, req)
	if err != nil {
		errorHandler.Error(err, w, "authentication failed: ", http.StatusInternalServerError)
		return
	}

	events, err := h.M.GetUpcomingEvents()
	if err != nil {
		errorHandler.Error(err, w, "can't get events: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   events,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("got all upcoming events")
}
