package models

import (
	"database/sql"
	"time"

	"github.com/commandwncos/api-booking/command/private/database"
)

type Event struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" binding:"required"`
	Description string    `db:"description" json:"description" binding:"required"`
	Location    string    `db:"location" json:"location" binding:"required"`
	DateTime    time.Time `db:"datetime" json:"datetime" binding:"required"`
	UserID      int64     `db:"user_id" json:"user_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

func (e *Event) Save(db *sql.DB) error {
	query := `
		INSERT INTO events (name, description, location, datetime, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	return db.QueryRow(
		query,
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.UserID,
	).Scan(&e.ID, &e.CreatedAt)
}

func GetAllEvents(db *sql.DB) ([]Event, error) {
	query := `
		SELECT 
			id, name, description, location, datetime, user_id, created_at
		FROM 
			events
		ORDER BY 
			datetime ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Description,
			&e.Location,
			&e.DateTime,
			&e.UserID,
			&e.CreatedAt,
		); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func GetEventsByUser(db *sql.DB, userID int64) ([]Event, error) {
	query := `
		SELECT 
			id, name, description, location, datetime, user_id, created_at
		FROM 
			events
		WHERE 
			user_id = $1
		ORDER BY datetime ASC
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Description,
			&e.Location,
			&e.DateTime,
			&e.UserID,
			&e.CreatedAt,
		); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func GetEventById(db *sql.DB, eventID, userID int64) (*Event, error) {
	query := `
		SELECT 
			id, name, description, location, datetime, user_id, created_at
		FROM 
			events
		WHERE 
			id = $1 AND user_id = $2
	`

	row := db.QueryRow(query, eventID, userID)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
		&event.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update(db *sql.DB) error {
	query := `
		UPDATE events
		SET name = $1,
		    description = $2,
		    location = $3,
		    datetime = $4
		WHERE id = $5 AND user_id = $6
	`

	result, err := db.Exec(
		query,
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.ID,
		e.UserID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (e *Event) UpdateEventById() error {
	query := `UPDATE events SET name = $1, description = $2, location = $3, datetime = $4 WHERE id = $5;`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) DeleteEventById(id int64) error {
	result, err := database.DB.Exec(
		"DELETE FROM events WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteEvent(db *sql.DB, eventID, userID int64) error {
	result, err := db.Exec(
		"DELETE FROM events WHERE id = $1 AND user_id = $2",
		eventID,
		userID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
