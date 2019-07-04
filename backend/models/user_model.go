package models

import (
	"database/sql"
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
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
		log.Println("password is correct")
		return user, err
	}
	return User{}, errors.New("incorect password")
}

func (model UserModel) CreateUser() {

}

func (model UserModel) DeleteUser() {

}

func (model UserModel) GetAllUsers() {

}

func (model UserModel) Logout() {

}
