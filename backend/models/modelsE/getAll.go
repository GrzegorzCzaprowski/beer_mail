package modelsE

func (model EventModel) GetAllEvents() ([]Event, error) {
	rows, err := model.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	events, err := model.scanRows(rows)
	if err != nil {
		return nil, err
	}
	return events, err
}
