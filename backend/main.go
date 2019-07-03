package main

import (
	"database/sql"
	"log"
	"net/http"

	user "github.com/GrzegorzCzaprowski/beer_mail/backend/handlers"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()

	userModel := models.UserModel{DB: db}
	userHandler := user.UserHandler{M: userModel}
	router.POST("user/add", userHandler.Post)

	log.Fatal(http.ListenAndServe(":8000", router))
}
