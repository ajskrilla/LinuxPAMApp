package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"log"
)

var DB *sql.DB

func Init(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		return err
	}

	log.Println("Database initialized.")
	return nil
}

