package handlers

import (
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h EventHandler) Events(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := authorization.UserAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusInternalServerError)
		return
	}

	events, err := h.M.GetAllEvents()
	if err != nil {
		error_handler.Error(err, w, "error with database: ", http.StatusInternalServerError)
		return
	}

	//guests, err := h.M.GetAllGuestsFromDB()

	res := response.Resp{
		Status: "succes",
		Data:   events,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("got all events")
}
