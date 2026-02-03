package models

import (
	"database/sql"
	"time"

	"github.com/commandwncos/api-booking/command/utils"
)

type User struct {
	ID           int64     `db:"id" json:"id"`
	Email        string    `db:"email" json:"email" binding:"required"`
	Password     string    `json:"password,omitempty" binding:"required"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

func (u *User) Save(db *sql.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	return db.QueryRow(
		query,
		u.Email,
		string(hashedPassword),
	).Scan(&u.ID, &u.CreatedAt)
}

func (u *User) ValidateCredentials(db *sql.DB) (bool, error) {
	query := `
		SELECT id, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	err := db.QueryRow(query, u.Email).
		Scan(&u.ID, &u.PasswordHash, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	check := utils.CheckPassword(u.Password, u.PasswordHash)
	u.Password = ""
	return check, nil
}
