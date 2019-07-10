package authorization

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ConfirmTokenAuthentication(w http.ResponseWriter, req *http.Request, params httprouter.Params) error {
	// var tokenString string

	// tokenString = params.ByName("token")

	// claims := &models.ClaimsC{}

	// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return models.JwtKey, nil
	// })
	// if !tkn.Valid {
	// 	return errors.New("token isn't valid")
	// }
	// if err != nil {
	// 	return err
	// }
	// claims.UserID

	return nil
}
