package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsE"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//Post event
func (h EventHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := authorization.UserTokenAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusInternalServerError)
		return
	}

	event := modelsE.Event{}
	err = json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		error_handler.Error(err, w, "error with decoding event from json: ", http.StatusInternalServerError)
		return
	}

	event.IDcreator = id
	eventID, err := h.M.InsertEventIntoDB(event)
	if err != nil {
		error_handler.Error(err, w, "error with inserting event to database: ", http.StatusInternalServerError)
		return
	}
	event.ID = eventID

	user, err := h.M.GetCreator(id)
	if err != nil {
		error_handler.Error(err, w, "error with getting event's creator from database: ", http.StatusInternalServerError)
		return
	}
	err = h.M.SendMailsToAllUsers(event, user)
	if err != nil {
		error_handler.Error(err, w, "error with sending emails to users: ", http.StatusInternalServerError)
		return
	}
	log.Info("mails sended")

	res := response.Resp{
		Status: "succes",
		Data:   event,
	}
	response.Writer(w, res, http.StatusOK) //nie zwraca id
	log.Info("created ", event.Name)

	///////////////

}
