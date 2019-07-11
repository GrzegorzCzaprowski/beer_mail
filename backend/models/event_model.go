package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/lib/pq"
	"gopkg.in/gomail.v2"
)

type Event struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IDcreator int    `json:"idcreator,omitempty"`
	Date      string `json:"date,omitempty"`
	Place     string `json:"place,omitempty"`
}

type EventModel struct {
	DB *sql.DB
}

type ClaimsC struct {
	UserID  int
	EventID int
	jwt.StandardClaims
}

func (model EventModel) InsertEventIntoDB(event Event) (int, error) {
	row := model.DB.QueryRow("INSERT INTO events(name, id_users, date, place) VALUES($1, $2, $3, $4) RETURNING id", event.Name, event.IDcreator, event.Date, event.Place)
	err := row.Scan(&event.ID)
	return event.ID, err
}

func (model EventModel) InsertGuestIntoDB(eventID, userID int) error {
	_, err := model.DB.Exec("INSERT INTO guests(id_events, id_users) VALUES($1, $2)", eventID, userID)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return err
		}
	}

	return err
}

func (model EventModel) SendMailsToAllUsers(event Event, creator User) error {
	rows, err := model.DB.Query("SELECT * FROM users")
	if err != nil {
		return err
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.IsAdmin)
		if err != nil {
			return err
		}

		if user.ID == creator.ID {
			continue
		}
		err = model.InsertGuestIntoDB(event.ID, user.ID)
		if err != nil {
			return err
		}

		m := createMessage(event, user, creator)
		d := gomail.NewDialer("smtp.googlemail.com", 465, "gespiwko@gmail.com", "6_#{aJ't*kd,") //TODO: ZMIENIC ADRES WYSYLAJACY NA FLAGE
		err = d.DialAndSend(m)
		if err != nil {
			panic(err)
		}
		log.Info("email sended to ", user.Email)

	}
	return rows.Err()
}

func (model EventModel) GetCreator(id int) (User, error) {
	var user User
	row := model.DB.QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Surname, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (model EventModel) GetAllEventsFromDB() ([]Event, error) {
	var events []Event
	rows, err := model.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.IDcreator, &event.Date, &event.Place)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	return events, rows.Err()
}

func (model EventModel) DeleteEventFromDB(id int) error {
	res, err := model.DB.Exec("DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return err
	}

	numberOfRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numberOfRows < 1 {
		return fmt.Errorf("event with id %d dont exists", id)
	}
	return err
}

func (model EventModel) GetEvent(id int) (Event, error) {
	var event Event
	row := model.DB.QueryRow("SELECT * FROM events WHERE id=$1", id)
	err := row.Scan(&event.ID, &event.Name, &event.IDcreator, &event.Date, &event.Place)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (model EventModel) ConfirmEventForUser(eventID, userID int) error {
	res, err := model.DB.Exec("UPDATE guests SET confirm = true WHERE id_events=$1 AND id_users=$2", eventID, userID)
	if err != nil {
		return err
	}

	numberOfRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numberOfRows < 1 {
		return fmt.Errorf("event with id %d dont exists", eventID)
	}
	return err
}

func createMessage(event Event, user, creator User) *gomail.Message {
	expirationTime := time.Now().Add(480 * time.Minute)

	claims := &ClaimsC{
		UserID:  user.ID,
		EventID: event.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {

		return nil
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "gespiwko@gmail.com") //TODO: ZMIENIC ADRES WYSYLAJACY NA FLAGE
	m.SetHeader("To", user.Email)
	m.SetAddressHeader("Cc", user.Email, user.Name+" "+user.Surname)
	m.SetHeader("Subject", event.Name)

	m.SetBody("text/html", "at: "+event.Date+", "+creator.Name+" invites you for beer in "+event.Place+"\n http://localhost:8000/event/confirm/"+tokenString)
	return m
}
