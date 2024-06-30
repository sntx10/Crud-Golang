package models

import (
	"database/sql"
	"log"
)

func ConnectDatabase(db *sql.DB) {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS articles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL
    );`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
