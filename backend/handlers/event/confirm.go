package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsE"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/dgrijalva/jwt-go"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h EventHandler) ConfirmLink(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var tokenString string
	tokenString = params.ByName("token")
	claims, err := claims(tokenString)
	if err != nil {
		log.Error("error with token: ", err)
		return
	}

	err = h.M.ConfirmEvent(claims.EventID, claims.UserID, claims.Confirm)
	if err != nil {
		log.Error("error with confirming user for this event: ", err)
		return
	}
	event, err := h.M.GetEvent(claims.EventID)
	if err != nil {
		log.Error("error with geting event: ", err)
		return
	}
	user, err := h.M.GetUser(claims.UserID)
	if err != nil {
		log.Error("error with geting user: ", err)
		return
	}
	if claims.Confirm {
		log.Info(fmt.Sprintf("User %s confirmed his participation in event %s", user.Name, event.Name))
	} else {
		log.Info(fmt.Sprintf("User %s denied his participation in event %s", user.Name, event.Name))
	}
}

func claims(tokenString string) (*modelsE.ClaimsC, error) {
	claims := &modelsE.ClaimsC{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return modelsU.JwtKey, nil
	})
	if !tkn.Valid {
		return claims, errors.New("token isn't valid")
	}
	if err != nil {
		return claims, err
	}
	return claims, nil
}
