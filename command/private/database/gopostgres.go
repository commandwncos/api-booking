package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() *sql.DB {
	_ = godotenv.Load()

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Panic("Could not open database connection:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	log.Println("✅ Connected to PostgreSQL")
	createTables()
	return DB
}

func createTables() {

	createUserTable := `
	CREATE EXTENSION IF NOT EXISTS citext;

	CREATE TABLE IF NOT EXISTS users (
		id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		email CITEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	`
	if _, err := DB.Exec(createUserTable); err != nil {
		log.Panic("Could not create users table:", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMPTZ NOT NULL,
		user_id BIGINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

		CONSTRAINT fk_events_user
			FOREIGN KEY (user_id)
			REFERENCES users(id)
			ON DELETE CASCADE
	);
	`
	if _, err := DB.Exec(createEventsTable); err != nil {
		log.Panic("Could not create events table:", err)
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT NOT NULL,
		event_id BIGINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

		CONSTRAINT fk_registrations_user
			FOREIGN KEY (user_id)
			REFERENCES users(id)
			ON DELETE CASCADE,

		CONSTRAINT fk_registrations_event
			FOREIGN KEY (event_id)
			REFERENCES events(id)
			ON DELETE CASCADE,

		CONSTRAINT uq_user_event UNIQUE (user_id, event_id)
	);
	`
	if _, err := DB.Exec(createRegistrationsTable); err != nil {
		log.Panic("Could not create registrations table:", err)
	}

	createIndexes()
}

func createIndexes() {

	createIndexes := `
	CREATE INDEX IF NOT EXISTS idx_events_user_id_id
		ON events(user_id, id);

	CREATE INDEX IF NOT EXISTS idx_registrations_user_id
		ON registrations(user_id);

	CREATE INDEX IF NOT EXISTS idx_registrations_event_id
		ON registrations(event_id);
	`

	if _, err := DB.Exec(createIndexes); err != nil {
		log.Panic("Could not create indexes:", err)
	}

	log.Println("📌 Database tables and indexes ensured")
}
