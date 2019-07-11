package modelsE

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/dgrijalva/jwt-go"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type Event struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	IDcreator int     `json:"idcreator,omitempty"`
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

func (model EventModel) SendMailsToAllUsers(event Event, creator modelsU.User) error {
	rows, err := model.DB.Query("SELECT * FROM users")
	if err != nil {
		return err
	}

	for rows.Next() {
		user := modelsU.User{}
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

func (model EventModel) GetCreator(id int) (modelsU.User, error) {
	var user modelsU.User
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
		//	//	//	//
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
		event.Guests = guests
		///////
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

func (model EventModel) ConfirmEventForUser(eventID, userID int, confirm bool) error {
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

func createMessage(event Event, user, creator modelsU.User) *gomail.Message {
	expirationTime := time.Now().Add(480 * time.Minute)

	claims := &ClaimsC{
		UserID:  user.ID,
		EventID: event.ID,
		Confirm: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenStringConfirm, err := token.SignedString(modelsU.JwtKey)
	if err != nil {
		return nil
	}

	claims.Confirm = false

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenStringDeny, err := token.SignedString(modelsU.JwtKey)
	if err != nil {
		return nil
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "gespiwko@gmail.com") //TODO: ZMIENIC ADRES WYSYLAJACY NA FLAGE
	m.SetHeader("To", user.Email)
	m.SetAddressHeader("Cc", user.Email, user.Name+" "+user.Surname)
	m.SetHeader("Subject", event.Name)

	m.SetBody("text/html", "at: "+event.Date+", "+creator.Name+" invites you for beer in "+event.Place+
		"\n if you want to confirm click this link: http://localhost:8000/event/confirm/"+tokenStringConfirm+
		"\n if you want to deny click this link: http://localhost:8000/event/confirm/"+tokenStringDeny)
	return m
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
