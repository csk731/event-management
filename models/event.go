package models

import (
	"chaitanyaallu.dev/event-management/db"
)

type Event struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	DateTime    string `json:"dateTime" binding:"required"`
	UserID      int    `json:"userId"`
}

func (event *Event) Save() error {
	query := `INSERT INTO EVENTS (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)`
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()
	result, err := stmnt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}
	eventID, err := result.LastInsertId()
	event.ID = eventID
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM EVENTS`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var events = []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	defer rows.Close()
	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT * FROM EVENTS WHERE id = ?`
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmnt.Close()
	row := stmnt.QueryRow(id)
	var event Event
	err = row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `UPDATE EVENTS 
				SET name = ?, description = ?, location = ?, dateTime = ? 
				WHERE id = ?`
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()
	result, err := stmnt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}

func (event *Event) DeleteEvent() error {
	query := `DELETE FROM EVENTS WHERE id = ?`
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()
	_, err = stmnt.Exec(event.ID)
	if err != nil {
		return err
	}
	return nil
}
