package repo

import (
	"context"

	"github.com/Moritisimor/shawty/internal/models"
)

func (repo *URLAliasRepo) GetAllAliases(ctx context.Context) ([]models.URLAlias, error) {
	acc := []models.URLAlias{}
	rows, err := repo.db.QueryContext(ctx, "SELECT ID, Alias, URL, DeleteAt FROM URLAliases")
	if err != nil {
		return acc, err
	}

	for rows.Next() {
		var id uint
		var alias, url string
		var deleteAt int64

		if err := rows.Scan(&id, &alias, &url, &deleteAt); err != nil {
			return acc, err
		}

		acc = append(acc, models.URLAlias{
			ID:       id,
			Alias:    alias,
			URL:      url,
			DeleteAt: deleteAt,
		})
	}

	return acc, rows.Err()
}
