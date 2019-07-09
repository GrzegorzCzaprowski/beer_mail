package authorization

import (
	"errors"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"
)

func UserTokenAuthentication(w http.ResponseWriter, req *http.Request) (int, error) {
	tokenstring := req.Header.Get("token")

	var id int

	if len(tokenstring) == 0 {
		return id, errors.New("token dont exists")
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !tkn.Valid {
		return id, errors.New("token isn't valid")
	}
	if err != nil {
		return id, err
	}

	id = claims.ID

	return id, nil
}
