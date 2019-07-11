package handlers

import (
	"net/http"
	"strconv"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h EventHandler) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	_, err := authorization.UserTokenAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		error_handler.Error(err, w, "can't parse id to int ", http.StatusInternalServerError)
		return
	}

	event, err := h.M.GetEvent(id)
	if err != nil {
		error_handler.Error(err, w, "can't get event: ", http.StatusInternalServerError)
		return
	}

	err = h.M.DeleteEventFromDB(id)
	if err != nil {
		error_handler.Error(err, w, "can't delete event: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   event,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("event with id ", id, " was deleted sucesfully")

	log.Info("event ", event.Name, " deleted")
}
