package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsE"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//Post event
func (h EventHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := authorization.UserAuthentication(w, req)
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

	event.IDcreator = strconv.Itoa(id)
	eventID, err := h.M.InsertEvent(event)
	if err != nil {
		error_handler.Error(err, w, "error with inserting event to database: ", http.StatusInternalServerError)
		return
	}
	event.ID = eventID
	creator, err := h.M.GetUser(id)
	if err != nil {
		error_handler.Error(err, w, "error with getting creator's event from database: ", http.StatusInternalServerError)
		return
	}
	err = h.M.SendMails(event, creator)
	if err != nil {
		error_handler.Error(err, w, "error with sending emails to users: ", http.StatusInternalServerError)
		return
	}
	log.Info("mails sended")
	event.IDcreator = creator.Name + " " + creator.Surname
	res := response.Resp{
		Status: "succes",
		Data:   event,
	}
	response.Writer(w, res, http.StatusOK) //nie zwraca id
	log.Info("created ", event.Name)

	///////////////

}
