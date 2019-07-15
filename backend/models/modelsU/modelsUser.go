package modelsU

import (
	"database/sql"

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
