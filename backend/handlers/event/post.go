package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//Post event
func (h EventHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := authorization.UserTokenAuthentication(w, req)
	if err != nil {
		log.Error("authentication failed: ", err)
		w.WriteHeader(500)
		return
	}

	event := models.Event{}
	err = json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		log.Error("error with decoding user from json: ", err)
		w.WriteHeader(500)
		return
	}

	event.IDcreator = id
	err = h.M.InsertEventIntoDB(event)
	if err != nil {
		log.Error("error with inserting user to database: ", err)
		w.WriteHeader(500)
		return
	}
	log.Info("created ", event.Name)
}
