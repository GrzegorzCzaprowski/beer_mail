package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (h UserHandler) Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
		Path:   "/",
	})

	log.Info("User logout sucesfully")
}
