package modelsE

import (
	"database/sql"
	"fmt"

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

func (model EventModel) DeleteEvent(id int) error {
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

func (model EventModel) ConfirmEvent(eventID, userID int, confirm bool) error {
	res, err := model.DB.Exec("UPDATE guests SET confirm =$1 WHERE id_events=$2 AND id_users=$3", confirm, eventID, userID)
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
