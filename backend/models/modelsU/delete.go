package modelsU

import "fmt"

func (model UserModel) DeleteUser(id int) error {
	res, err := model.DB.Exec("DELETE FROM guests WHERE id_users=$1", id)
	if err != nil {
		return err
	}

	res, err = model.DB.Exec("DELETE FROM users WHERE id=$1", id)
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
