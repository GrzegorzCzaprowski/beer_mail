package authorization

import (
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"
)

func AdminTokenAuthentication(w http.ResponseWriter, req *http.Request) (bool, error) {
	cookie, err := req.Cookie("token")
	if err != nil {
		return false, err
	}

	tknStr := cookie.Value

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !tkn.Valid {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if claims.IsAdmin {
		return true, nil
	}
	return false, nil

}
