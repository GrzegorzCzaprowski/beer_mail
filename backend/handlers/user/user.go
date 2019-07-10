package handlers

import (
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//User returns user info
func (h UserHandler) User(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := authorization.UserTokenAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusInternalServerError)
		return
	}

	user, err := h.M.GetUser(id)
	if err != nil {
		error_handler.Error(err, w, "error with decoding user from json: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   user,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("got an user")
}
