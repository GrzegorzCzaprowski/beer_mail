package modelsE

import "fmt"

func (model EventModel) DeleteEvent(id int) error {
	res, err := model.DB.Exec("DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return err
	}

	numberOfRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numberOfRows < 1 {
		return fmt.Errorf("event with id %d dont exists", id)
	}
	return err
}
