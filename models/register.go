package models

import (
	"database/sql"
	"errors"
	"strings"
)

var ErrAlreadyRegistered = errors.New("already registered")

func RegisterUserForEvent(db *sql.DB, userID, eventID int64) error {
	query := `
	INSERT INTO registrations (user_id, event_id)
	VALUES ($1, $2)
	`

	_, err := db.Exec(query, userID, eventID)
	if err != nil {
		if strings.Contains(err.Error(), "uq_user_event") {
			return ErrAlreadyRegistered
		}
		return err
	}
	return nil
}

func CancelRegistration(db *sql.DB, userID, eventID int64) (int64, error) {
	query := `
	DELETE FROM registrations
	WHERE user_id = $1 AND event_id = $2
	`

	result, err := db.Exec(query, userID, eventID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func IsEventOwner(db *sql.DB, eventID, userID int64) (bool, error) {
	var exists bool
	query := `
	SELECT EXISTS (
		SELECT 1 FROM events
		WHERE id = $1 AND user_id = $2
	)
	`
	err := db.QueryRow(query, eventID, userID).Scan(&exists)
	return exists, err
}
