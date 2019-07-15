package modelsE

import (
	"strings"
	"time"
)

func (model EventModel) GetUpcomingEvents() ([]Event, error) {
	date := time.Now()
	d := strings.Split(date.Format(time.RFC3339), "+")
	d[0] = strings.ReplaceAll(d[0], "T", " ")

	rows, err := model.DB.Query("SELECT * FROM events WHERE date >= to_timestamp($1, 'yyyy-mm-dd hh24:mi:ss')", d[0])
	if err != nil {
		return nil, err
	}

	events, err := model.scanRows(rows)
	if err != nil {
		return nil, err
	}
	return events, err
}
