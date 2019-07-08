package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h UserHandler) Get(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ok, err := authorization.AdminTokenAuthentication(w, req)
	if err != nil {
		w.WriteHeader(500)
		log.Error("authentication failed: ", err)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		log.Warn("you are not an admin")
		return
	}

	users, err := h.M.GetAllUsersFromDB()
	if err != nil {
		log.Error("error with decoding user from json: ", err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Error("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
	log.Info("got all users")
}
