package error_handler

import (
	"errors"
	"net/http"

	l "github.com/sirupsen/logrus"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
)

var NotAdmin = errors.New("you are not an admin")

func Error(err error, w http.ResponseWriter, log string, httpstatus int) {
	res := response.Resp{
		Status: "error",
		Data:   err.Error(),
	}
	response.Writer(w, res, httpstatus)
	l.Error(log, err)
}
