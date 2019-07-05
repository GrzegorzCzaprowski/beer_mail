package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h UserHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ok, err := authorization.AdminTokenAuthentication(w, req)
	if err != nil {
		log.Error("authentication failed: ", err)
		return
	}
	if !ok {
		log.Warn("you are not an admin")
		return
	}

	user := models.User{}
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Error("error with decoding user from json: ", err)
		w.WriteHeader(500)
		return
	}

	err = h.M.InsertUserIntoDB(user)
	if err != nil {
		log.Error("error with inserting user to database: ", err)
		w.WriteHeader(500)
		return
	}
	log.Info("created ", user.Email)
}
