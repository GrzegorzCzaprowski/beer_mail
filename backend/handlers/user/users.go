package handlers

import (
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//Users returns all users
func (h UserHandler) Users(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := authorization.AdminTokenAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusUnauthorized)
		return
	}

	users, err := h.M.GetAllUsersFromDB()
	if err != nil {
		error_handler.Error(err, w, "error with getting users from database: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   users,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("got all users")
}
