package repo

import (
	"context"
	"time"

	"github.com/Moritisimor/shawty/internal/models"
)

func (repo *URLAliasRepo) PostURLAlias(
	urlAlias models.URLAliasDTO,
	ctx context.Context,
) (int64, error) {
	rn := time.Now().Unix()
	results, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO URLAliases (Alias, URL, DeleteAt) VALUES (?, ?, ?)",
		urlAlias.Alias, urlAlias.URL, rn+(repo.deleteAfterHours*3600),
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
