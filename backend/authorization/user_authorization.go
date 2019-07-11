package authorization

import (
	"errors"
	"net/http"
	"strings"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/dgrijalva/jwt-go"
)

func UserTokenAuthentication(w http.ResponseWriter, req *http.Request) (int, error) {
	header := req.Header.Get("Authorization")
	var id int

	AuthArr := strings.Split(header, " ")
	var tokenString string
	if len(AuthArr) == 2 {
		tokenString = AuthArr[1]
	} else {
		return id, errors.New("token isn't valid")
	}

	claims := &modelsU.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return modelsU.JwtKey, nil
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
