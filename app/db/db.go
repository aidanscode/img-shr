package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrMigrationFailure	= errors.New("failed to initialize database")
)

type Database struct {
	Db *sql.DB
}

func (d *Database) Migrate() error {
	createPostsTable := `
		CREATE TABLE IF NOT EXISTS posts(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			img_path TEXT,
			created_at TIMESTAMP NOT NULL
		)
	`
	_, createPostsError := d.Db.Exec(createPostsTable)
	if createPostsError != nil {
		return ErrMigrationFailure
	}

	return nil
}

func NewDatabase(path string) (*Database, error) {
	sqlite, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	database := Database{Db: sqlite}
	return &database, nil
}
