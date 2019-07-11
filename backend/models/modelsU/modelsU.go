package modelsU

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"isadmin,omitempty"`
}

var JwtKey = []byte("nothing to see here goy")

type Claims struct {
	ID      int  `json:"id"`
	IsAdmin bool `json:"isadmin"`
	jwt.StandardClaims
}

type UserModel struct {
	DB *sql.DB
}

func (model UserModel) FindUserInDB(user User) (User, error) {
	password := user.Password
	row := model.DB.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Surname, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	if password == user.Password {
		return user, err
	}
	return User{}, errors.New("incorect password")
}

func (model UserModel) InsertUserIntoDB(user User) error {
	_, err := model.DB.Exec("INSERT INTO users(name, surname, email, password, admin) VALUES($1, $2, $3, $4, $5)", user.Name, user.Surname, user.Email, user.Password, user.IsAdmin)
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

func (model UserModel) DeleteUserFromDB(id int) error {
	res, err := model.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	numberOfRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numberOfRows < 1 {
		return fmt.Errorf("user with id %d dont exists", id)
	}
	return err
}

func (model UserModel) GetAllUsersFromDB() ([]User, error) {
	var users []User
	rows, err := model.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.IsAdmin)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (model UserModel) GetUser(id int) (User, error) {
	var user User
	row := model.DB.QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Surname, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	return user, nil
}
