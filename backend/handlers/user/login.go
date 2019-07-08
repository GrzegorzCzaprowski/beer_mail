package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/error_handler"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		error_handler.Error(err, w, "error with decoding user from json: ", http.StatusInternalServerError)
		return
	}

	user, err = h.M.FindUserInDB(user)
	if err != nil {
		error_handler.Error(err, w, "cant find this user: ", http.StatusInternalServerError)
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
		error_handler.Error(err, w, "error with creating token: ", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/", //USTAWIA COOKIE NA DOMYSLNY PATH /, WIEC COOKIE JEST DOSTEPNE WSZEDZIE KURWA
	})

	res := response.Resp{
		Status: "succes",
		Data:   user,
	}
	response.Writer(w, res, http.StatusOK)
	log.Info("You loged correctly")
}
