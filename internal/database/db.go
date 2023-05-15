package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func InitializeDatabase(dbFile string, schemaFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	schema, err := os.ReadFile(schemaFile)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, err
	}

	return db, nil
}
