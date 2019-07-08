package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h EventHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ok, err := authorization.UserTokenAuthentication(w, req)
	if err != nil {
		log.Error("authentication failed: ", err)
		w.WriteHeader(500)
		return
	}
	if !ok {
		log.Warn("something wrong with session") //TODO: cos tutaj nie teges, prawdopodobnie usunąć tego warna
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

	err = h.M.InsertEventIntoDB(event)
	if err != nil {
		log.Error("error with inserting user to database: ", err)
		w.WriteHeader(500)
		return
	}
	// log.Info("created ", user.Email)
}
