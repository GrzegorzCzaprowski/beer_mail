package main

// insert into users (name, surname, email, password, admin) values ('test', 'test', 'test@email.com', 'test', false);
// insert into users (name, surname, email, password, admin) values ('admin', 'admin', 'admin@email.com', 'admin', true);

import (
	"database/sql"
	"net/http"

	log "github.com/sirupsen/logrus"

	event "github.com/GrzegorzCzaprowski/beer_mail/backend/handlers/event"
	user "github.com/GrzegorzCzaprowski/beer_mail/backend/handlers/user"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsE"
	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()
	log.Info("program is running")

	userModel := modelsU.UserModel{DB: db}
	userHandler := user.UserHandler{M: userModel}
	router.POST("/user/post", userHandler.Post)
	router.POST("/user/login", userHandler.Login)
	router.GET("/user/get", userHandler.Users)
	router.DELETE("/user/delete/:id", userHandler.Delete)
	router.POST("/user/logout", userHandler.Logout)
	router.GET("/user", userHandler.User)

	eventModel := modelsE.EventModel{DB: db}
	eventHandler := event.EventHandler{M: eventModel}
	router.POST("/event/post", eventHandler.Post)
	router.GET("/event/get", eventHandler.Events)
	router.DELETE("/event/delete/:id", eventHandler.Delete)
	router.GET("/event/confirm/:token", eventHandler.ConfirmLink)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		// Enable Debugging for testing, consider disabling in production
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
