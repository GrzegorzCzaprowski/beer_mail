package modelsE

import (
	"database/sql"
	"strconv"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/dgrijalva/jwt-go"
)

type Event struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	IDcreator string  `json:"idcreator,omitempty"`
	Date      string  `json:"date,omitempty"`
	Place     string  `json:"place,omitempty"`
	Guests    []Guest `json:"guests,omitempty"`
}

type EventModel struct {
	DB *sql.DB
}

type ClaimsC struct {
	UserID  int
	EventID int
	Confirm bool
	jwt.StandardClaims
}

type Guest struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Email   string `json:"email,omitempty"`
	Confirm string `json:"confirm,omitempty"`
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

func (model EventModel) GetUser(id int) (modelsU.User, error) {
	var user modelsU.User
	row := model.DB.QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Surname, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (model EventModel) scanRows(rows *sql.Rows) ([]Event, error) {
	var events []Event
	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.IDcreator, &event.Date, &event.Place)
		if err != nil {
			return events, err
		}
		//	//	//	//
		guests, err := model.getGuests(event)
		if err != nil {
			return nil, err
		}
		event.Guests = guests

		id, err := strconv.Atoi(event.IDcreator)
		creator, err := model.GetUser(id)
		event.IDcreator = creator.Name + " " + creator.Surname

		///////
		events = append(events, event)
	}
	return events, rows.Err()
}

func (model EventModel) getGuests(event Event) ([]Guest, error) {
	var guests []Guest
	rows2, err := model.DB.Query("SELECT * FROM guests WHERE id_events=$1", event.ID)
	for rows2.Next() {
		guest := Guest{}
		var bla string
		var bla2 string
		var guestID int
		var tof string
		_ = rows2.Scan(&bla, &bla2, &guestID, &tof)
		if err != nil {
			return nil, err
		}
		guest.Confirm = tof

		user, err := model.GetUser(guestID)
		if err != nil {
			return nil, err
		}
		guest.Name = user.Name
		guest.Surname = user.Surname
		guest.Email = user.Email

		guests = append(guests, guest)
	}
	if err != nil {
		return nil, err
	}
	return guests, nil
}
