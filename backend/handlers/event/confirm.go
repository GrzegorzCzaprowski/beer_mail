package handlers

import (
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h EventHandler) Confirm(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var tokenString string

	tokenString = params.ByName("token")

	claims := &models.ClaimsC{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !tkn.Valid {
		return //errors.New("token isn't valid")
	}
	if err != nil {
		return //err
	}

	h.M.ConfirmEventForUser(claims.EventID, claims.UserID)
	var user models.User
	var event models.Event
	log.Info("User " + user.Name + " confirmed his participation in event " + event.Name)
}
