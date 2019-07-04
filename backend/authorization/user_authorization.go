package authorization

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"
)

func UserTokenAuthentication(w http.ResponseWriter, req *http.Request) (bool, error) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("to ")
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return false, err
		}
		// For any other type of error, return a bad request status
		log.Println("czy to ")
		w.WriteHeader(http.StatusBadRequest)
		return false, err
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &models.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false, nil
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return false, err
	}

	log.Println("welcome user")
	w.Write([]byte(fmt.Sprintf("Welcome admin")))
	return true, nil
}
