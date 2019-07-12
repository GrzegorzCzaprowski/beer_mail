package modelsE

import (
	"strconv"
)

func (model EventModel) GetAllEvents() ([]Event, error) {
	var events []Event
	rows, err := model.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.IDcreator, &event.Date, &event.Place)
		if err != nil {
			return events, err
		}
		//	//	//	//
		guests, err := model.getGuests(event)
		if err != nil {
			return nil, err
		}
		event.Guests = guests

		id, err := strconv.Atoi(event.IDcreator)
		creator, err := model.GetUser(id)
		event.IDcreator = creator.Name + " " + creator.Surname

		///////
		events = append(events, event)
	}
	return events, rows.Err()
}

func (model EventModel) getGuests(event Event) ([]Guest, error) {
	var guests []Guest
	rows2, err := model.DB.Query("SELECT * FROM guests WHERE id_events=$1", event.ID)
	for rows2.Next() {
		guest := Guest{}
		var bla string
		var bla2 string
		var guestID int
		var tof string
		_ = rows2.Scan(&bla, &bla2, &guestID, &tof)
		if err != nil {
			return nil, err
		}
		guest.Confirm = tof

		user, err := model.GetUser(guestID)
		if err != nil {
			return nil, err
		}
		guest.Name = user.Name
		guest.Surname = user.Surname
		guest.Email = user.Email

		guests = append(guests, guest)
	}
	if err != nil {
		return nil, err
	}
	return guests, nil
}
