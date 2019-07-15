package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/errorHandler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//Post posts user
func (h UserHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := authorization.AdminAuthentication(w, req)
	if err != nil {
		errorHandler.Error(err, w, "authentication failed: ", http.StatusInternalServerError)
		return
	}

	user := modelsU.User{}
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		errorHandler.Error(err, w, "error with decoding user from json: ", http.StatusInternalServerError)
		return
	}

	err = h.M.InsertUser(user)
	if err != nil {
		errorHandler.Error(err, w, "error with inserting user to database: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   user,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("created ", user.Email)
}
