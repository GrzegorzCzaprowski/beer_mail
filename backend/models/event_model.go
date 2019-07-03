package models

import "database/sql"

type Event struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Users []User `json:"users,omitempty"`
	Date  string `json:"date,omitempty"`
	Place string `json:"place,omitempty"`
}

type EventModel struct {
	DB *sql.DB
}

func (model EventModel) CreateEvent() {}
