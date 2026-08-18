package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/ncruces/go-sqlite3/driver"
)

type URLAliasRepo struct {
	db *sql.DB
}

func New(path string) (*URLAliasRepo, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS URLAliases (
			ID INTEGER PRIMARY KEY,
			Alias TEXT UNIQUE NOT NULL,
			URL TEXT NOT NULL
		)`,
	); err != nil {
		return nil, fmt.Errorf("Error while migrating: %s", err)
	}

	return &URLAliasRepo{
		db: db,
	}, nil
}
