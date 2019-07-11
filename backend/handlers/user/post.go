package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//Post posts user
func (h UserHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := authorization.AdminTokenAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusUnauthorized)
		return
	}

	user := models.User{}
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		error_handler.Error(err, w, "error with decoding user from json: ", http.StatusInternalServerError)
		return
	}

	err = h.M.InsertUserIntoDB(user)
	if err != nil {
		error_handler.Error(err, w, "error with inserting user to database: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   user,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("created ", user.Email)
}
