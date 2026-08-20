package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/Moritisimor/shawty/internal/models"
)

func (repo *URLAliasRepo) PostURLAlias(
	urlAlias models.URLAliasDTO,
	ctx context.Context,
) (int64, error) {
	var results sql.Result
	var err error

	if repo.deleteAfterHours > 0 {
		rn := time.Now().Unix()
		results, err = repo.db.ExecContext(
			ctx,
			"INSERT INTO URLAliases (Alias, URL, DeleteAt) VALUES (?, ?, ?)",
			urlAlias.Alias, urlAlias.URL, rn+(repo.deleteAfterHours*3600),
		)
	} else {
		results, err = repo.db.ExecContext(
			ctx,
			"INSERT INTO URLAliases (Alias, URL, DeleteAT) VALUES (?, ?, ?)",
			urlAlias.Alias, urlAlias.URL, nil,
		)
	}

	if err != nil {
		return 0, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
