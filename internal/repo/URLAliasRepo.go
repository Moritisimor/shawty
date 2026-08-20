package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/ncruces/go-sqlite3/driver"
)

type URLAliasRepo struct {
	db               *sql.DB
	deleteAfterHours int64
}

func New(path string, deleteAfterHours int64) (*URLAliasRepo, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS URLAliases (
			ID INTEGER PRIMARY KEY,
			Alias TEXT UNIQUE NOT NULL,
			URL TEXT NOT NULL,
			DeleteAt INTEGER
		)`,
	); err != nil {
		return nil, fmt.Errorf("Error while migrating: %s", err)
	}

	return &URLAliasRepo{
		db:               db,
		deleteAfterHours: deleteAfterHours,
	}, nil
}

func (u *URLAliasRepo) Close() error {
	return u.db.Close()
}
