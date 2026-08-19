package repo

import (
	"context"

	"github.com/Moritisimor/shawty/internal/models"
)

func (repo *URLAliasRepo) PostURLAlias(
	urlAlias models.URLAliasDTO, 
	ctx context.Context,
) (int64, error) {
	results, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO URLAliases (Alias, URL) VALUES (?, ?)",
		urlAlias.Alias, urlAlias.URL,
	)

	if err != nil {
		return 0, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
