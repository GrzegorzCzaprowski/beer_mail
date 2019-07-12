package modelsU

func (model UserModel) GetAllUsers() ([]User, error) {
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
