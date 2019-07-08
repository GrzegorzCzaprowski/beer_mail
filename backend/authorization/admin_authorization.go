package authorization

import (
	"errors"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"
)

func AdminTokenAuthentication(w http.ResponseWriter, req *http.Request) error {
	tokenstring := req.Header.Get("token")

	if len(tokenstring) == 0 {
		return errors.New("token dont exists")
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !claims.IsAdmin && !tkn.Valid {
		return errors.New("token isn't valid")
	}
	if !claims.IsAdmin {
		return errors.New("you are not an admin")
	}
	if err != nil {
		return err
	}

	return nil
}
