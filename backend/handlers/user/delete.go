package handlers

import (
	"net/http"
	"strconv"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"

	log "github.com/sirupsen/logrus"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/julienschmidt/httprouter"
)

//Delete it delete user
func (h UserHandler) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := authorization.AdminAuthentication(w, req)
	if err != nil {
		error_handler.Error(err, w, "authentication failed: ", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		error_handler.Error(err, w, "can't parse id to int ", http.StatusInternalServerError)
		return
	}

	user, err := h.M.GetUser(id)
	if err != nil {
		error_handler.Error(err, w, "can't get user: ", http.StatusInternalServerError)
		return
	}

	err = h.M.DeleteUser(id)
	if err != nil {
		error_handler.Error(err, w, "can't delete user: ", http.StatusInternalServerError)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   user,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("user with id ", id, " was deleted sucesfully")
}
