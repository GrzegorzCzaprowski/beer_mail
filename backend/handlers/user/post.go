package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ok, err := authorization.AdminTokenAuthentication(w, req)
	if err != nil {
		log.Println("authentication failed: ", err)
		return
	}
	if !ok {
		log.Println("you are not an admin")
		return
	}

	user := models.User{}
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("error with decoding user from json: ", err)
		w.WriteHeader(500)
		return
	}

	err = h.M.InsertUserIntoDB(user)
	if err != nil {
		log.Println("error with inserting user to database: ", err)
		w.WriteHeader(500)
		return
	}
	log.Println("created ", user.Name)
}
