package repo

import (
	"context"

	"github.com/Moritisimor/shawty/internal/models"
)

func (repo *URLAliasRepo) PostURLAlias(
	urlAlias models.URLAliasDTO, 
	ctx context.Context,
) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO URLAliases (Alias, URL) VALUES (?, ?)",
		urlAlias.Alias, urlAlias.URL,
	)

	return err
}
