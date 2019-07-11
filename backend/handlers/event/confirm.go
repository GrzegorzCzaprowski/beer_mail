package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
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

	h.M.ConfirmEventForUser(claims.EventID, claims.UserID, claims.Confirm)
	if err != nil {
		log.Error("error with confirming user for this event: ", err)
		return
	}

	if claims.Confirm {
		log.Info(fmt.Sprintf("User with ID %d confirmed his participation in event with ID %d", claims.UserID, claims.EventID))
	} else {
		log.Info(fmt.Sprintf("User with ID %d denied his participation in event with ID %d", claims.UserID, claims.EventID))
	}
}

func claims(tokenString string) (*models.ClaimsC, error) {
	claims := &models.ClaimsC{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !tkn.Valid {
		return claims, errors.New("token isn't valid")
	}
	if err != nil {
		return claims, err
	}
	return claims, nil
}
