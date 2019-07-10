package models

import (
	"database/sql"

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

func (model EventModel) InsertEventIntoDB(event Event) error {
	_, err := model.DB.Exec("INSERT INTO events(name, id_users, date, place) VALUES($1, $2, $3, $4)", event.Name, event.IDcreator, event.Date, event.Place)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return err
		}
	}

	if err != nil {
		return err
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
	return nil, nil
}

func createMessage(event Event, user, creator User) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "gespiwko@gmail.com") //TODO: ZMIENIC ADRES WYSYLAJACY NA FLAGE
	m.SetHeader("To", user.Email)
	m.SetAddressHeader("Cc", user.Email, user.Name+" "+user.Surname)
	m.SetHeader("Subject", event.Name)

	m.SetBody("text/html", "at: "+event.Date+", "+creator.Name+" invites you for beer in "+event.Place)
	return m
}
