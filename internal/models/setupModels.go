package models

import (
	"database/sql"
	"log"
)

func ConnectDatabase(db *sql.DB) {
	createArticlesTableQuery := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);`

	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := db.Exec(createArticlesTableQuery)
	if err != nil {
		log.Fatal("Failed to create articles table:", err)
	}

	_, err = db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}
}
