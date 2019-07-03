package main

import (
	"database/sql"
	"log"
	"net/http"

	event "github.com/GrzegorzCzaprowski/beer_mail/backend/handlers/event"
	user "github.com/GrzegorzCzaprowski/beer_mail/backend/handlers/user"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models"
	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
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
	router.POST("/user/add", userHandler.Post)

	eventModel := models.EventModel{DB: db}
	eventHandler := event.EventHandler{M: eventModel}
	router.POST("/event/add", eventHandler.Post)

	log.Fatal(http.ListenAndServe(":8000", router))
}
