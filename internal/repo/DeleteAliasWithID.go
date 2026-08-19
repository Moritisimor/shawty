package repo

import "context"

func (repo *URLAliasRepo) DeleteAliasWithID(id uint, ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM URLAliases WHERE id = ?", id)
	return err
}
