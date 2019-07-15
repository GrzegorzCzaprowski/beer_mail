package errorHandler

import (
	"net/http"

	l "github.com/sirupsen/logrus"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/response"
)

func Error(err error, w http.ResponseWriter, log string, httpstatus int) {
	res := response.Resp{
		Status: "error",
		Data:   err.Error(),
	}
	response.Writer(w, res, httpstatus)
	l.Error(log, err)
}
