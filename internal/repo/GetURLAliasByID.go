package repo

import (
	"context"

	"github.com/Moritisimor/shawty/internal/models"
)

func (u *URLAliasRepo) GetURLAliasByID(
	queryID uint,
	ctx context.Context,
) (models.URLAlias, error) {
	row := u.db.QueryRowContext(
		ctx,
		"SELECT ID, Alias, URL FROM URLAliases WHERE id = ?", queryID,
	)

	var id uint
	var alias, url string

	if err := row.Scan(&id, &alias, &url); err != nil {
		return models.URLAlias{}, err
	}

	return models.URLAlias{
		ID:    id,
		Alias: alias,
		URL:   url,
	}, nil
}
