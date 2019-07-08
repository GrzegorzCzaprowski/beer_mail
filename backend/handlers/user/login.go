package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Error("error with decoding user from json: ", err)
		w.WriteHeader(500)
		return
	}

	user, err = h.M.FindUserInDB(user)
	if err != nil {
		log.Error("can't find this user: ", err)
		w.WriteHeader(500)
		return
	}
	log.Info("password is correct")

	//lenghth of session for single user
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		ID:      user.ID,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(models.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/", //USTAWIA COOKIE NA DOMYSLNY PATH /, WIEC COOKIE JEST DOSTEPNE WSZEDZIE KURWA
	})
	log.Info("You loged correctly")
}
