package authorization

import (
	"errors"
	"net/http"
	"strings"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/dgrijalva/jwt-go"
)

func AdminTokenAuthentication(w http.ResponseWriter, req *http.Request) error {
	header := req.Header.Get("Authorization")

	AuthArr := strings.Split(header, " ")
	var tokenString string
	if len(AuthArr) == 2 {
		tokenString = AuthArr[1]
	} else {
		return errors.New("token isn't valid")
	}

	claims := &modelsU.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return modelsU.JwtKey, nil
	})
	if !tkn.Valid {
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
