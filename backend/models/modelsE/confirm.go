package modelsE

import "fmt"

func (model EventModel) ConfirmEvent(eventID, userID int, confirm bool) error {
	res, err := model.DB.Exec("UPDATE guests SET confirm =$1 WHERE id_events=$2 AND id_users=$3", confirm, eventID, userID)
	if err != nil {
		return err
	}

	numberOfRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numberOfRows < 1 {
		return fmt.Errorf("event with id %d dont exists", eventID)
	}
	return err
}
