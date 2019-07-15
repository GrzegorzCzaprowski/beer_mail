package modelsE

import (
	"time"

	"github.com/GrzegorzCzaprowski/beer_mail/backend/models/modelsU"
	"github.com/dgrijalva/jwt-go"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func (model EventModel) InsertEvent(event Event) (int, error) {
	row := model.DB.QueryRow("INSERT INTO events(name, id_users, date, place) VALUES($1, $2, $3, $4) RETURNING id",
		event.Name, event.IDcreator, event.Date, event.Place)
	err := row.Scan(&event.ID)
	return event.ID, err
}

func (model EventModel) SendMails(event Event, creator modelsU.User) error {
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
		err = model.insertGuest(event.ID, user.ID)
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

func (model EventModel) insertGuest(eventID, userID int) error {
	_, err := model.DB.Exec("INSERT INTO guests(id_events, id_users) VALUES($1, $2)", eventID, userID)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return err
		}
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
	tokenStringConfirm, err := token.SignedString(modelsU.JwtKey)
	if err != nil {
		return nil
	}

	claims.Confirm = false

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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
