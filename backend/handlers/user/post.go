package handlers

import (
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
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
	log.Println("admin robi adminowe rzeczy, np postuje uzytkownika")
}
