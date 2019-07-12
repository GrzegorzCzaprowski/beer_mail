package modelsU

import "github.com/lib/pq"

func (model UserModel) InsertUser(user User) error {
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
