package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeTable(db *sql.DB) error {
	q1 := `
		CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		name VARCHAR(255) NOT NULL,
		phone_number VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(50) NOT NULL CHECK (role IN ('customer', 'teller')) DEFAULT 'customer',
		email VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`

	q2 := `
		CREATE TABLE IF NOT EXISTS services (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		code TEXT UNIQUE,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`

	q3 := `
		CREATE TABLE IF NOT EXISTS queues (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		status VARCHAR(50) NOT NULL CHECK (status IN ('pending', 'processed', 'completed')) DEFAULT 'pending',
		queue_number varchar(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		service_id UUID REFERENCES services(id) ON DELETE cascade
	);`

	if _, err := db.Exec(q1); err != nil {
		log.Printf("initialize table users: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q2); err != nil {
		log.Printf("initialize table services: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q3); err != nil {
		log.Printf("initialize table queues: %s\n", err.Error())
		return err
	}

	return nil
}
