package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS students (
		id TEXT PRIMARY KEY,
		name TEXT
	);

	CREATE TABLE IF NOT EXISTS lessons (
		id TEXT PRIMARY KEY,
		name TEXT,
		date TEXT,
		body TEXT,
		student_id TEXT,
		FOREIGN KEY(student_id) REFERENCES students(id)
	);

	CREATE TABLE IF NOT EXISTS homework (
		id TEXT PRIMARY KEY,
		name TEXT,
		date TEXT,
		body TEXT,
		student_id TEXT,
		FOREIGN KEY(student_id) REFERENCES students(id)
	);
	`

	_, err = db.Exec(createTables)
	if err != nil {
		return nil, err
	}

	return db, nil
}
