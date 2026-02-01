package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("database connected successfully.")
	return db, nil
}
