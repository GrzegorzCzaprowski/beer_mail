package handlers

import (
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/authorization"
	"github.com/julienschmidt/httprouter"
)

func (h EventHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ok, err := authorization.UserTokenAuthentication(w, req)
	if err != nil {
		log.Println("authentication failed: ", err)
		return
	}
	if !ok {
		log.Println("you are not logged")
		return
	}

	log.Println("user robi userowe rzeczy, np zaklada nowy event")
}
