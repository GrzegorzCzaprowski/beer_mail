package modelsU

func (model UserModel) GetUser(id int) (User, error) {
	var user User
	row := model.DB.QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Surname, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	return user, nil
}
