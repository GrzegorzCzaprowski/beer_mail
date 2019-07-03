package models

import "database/sql"

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"isadmin,omitempty"`
}

type UserModel struct {
	DB *sql.DB
}

func (model UserModel) CreateUser() {

}

func (model UserModel) DeleteUser() {

}

func (model UserModel) GetAllUsers() {

}

func (model UserModel) Login() {

}

func (model UserModel) Logout() {

}
