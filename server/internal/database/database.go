package database

import "database/sql"

type Database struct {
	Db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{Db: db}
}

func (d *Database) Create() error {
	_, err := d.Db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			guid TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			start_date DATE NULL,
			end_date DATE NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
