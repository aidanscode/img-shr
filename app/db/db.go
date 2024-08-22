package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrMigrationFailure	= errors.New("failed to initialize database")
)

type Database interface {
	Migrate() error
}

type database struct {
	db *sql.DB
}

func (d *database) Migrate() error {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		)
	`
	_, createUsersError := d.db.Exec(createUsersTable)
	if createUsersError != nil {
		return ErrMigrationFailure
	}

	createPostsTable := `
		CREATE TABLE IF NOT EXISTS posts(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			author_id INTEGER,
			title TEXT NOT NULL,
			img_path TEXT NOT NULL UNIQUE,
			created_at TIMESTAMP NOT NULL,
			FOREIGN KEY (author_id) REFERENCES users(id)
		)
	`
	_, createPostsError := d.db.Exec(createPostsTable)
	if createPostsError != nil {
		return ErrMigrationFailure
	}

	return nil
}

func NewDatabase(path string) (Database, error) {
	sqlite, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	database := database{db: sqlite}
	return &database, nil
}
