package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Error("cant parse paramater to int: ", err)
		w.WriteHeader(500)
		return
	}

	user, err := h.M.GetUserBeforeDeletion(id)
	if err != nil {
		log.Error("can't get user: ", err)
		w.WriteHeader(500)
		return
	}

	err = h.M.DeleteUserFromDB(id)
	if err != nil {
		log.Error("can't delete user: ", err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Error("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
	log.Info("user with id ", id, " was deleted sucesfully")
}
