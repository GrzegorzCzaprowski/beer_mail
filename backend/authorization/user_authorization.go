package authorization

import (
	"errors"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"
)

func UserTokenAuthentication(w http.ResponseWriter, req *http.Request) error {
	tokenstring := req.Header.Get("token")

	if len(tokenstring) == 0 {
		return errors.New("token dont exists")
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !tkn.Valid {
		return errors.New("token isn't valid")
	}
	if err != nil {
		return err
	}

	return nil
}
