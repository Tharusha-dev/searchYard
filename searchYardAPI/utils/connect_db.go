package utils

import (
	"database/sql"
	"log"
)

func ConnectionToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "../searchYardAPI/db/main.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
