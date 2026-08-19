package repo

import "context"

func (r *URLAliasRepo) GetURLByAlias(
	alias string,
	ctx context.Context,
) (string, error) {
	row := r.db.QueryRowContext(
		ctx,
		"SELECT URL FROM URLAliases WHERE Alias = ?",
		alias,
	)

	var url string
	err := row.Scan(&url)
	return url, err
}
