package modelsU

import "errors"

func (model UserModel) FindUser(user User) (User, error) {
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
