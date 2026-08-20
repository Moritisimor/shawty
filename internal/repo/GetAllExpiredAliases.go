package repo

import (
	"context"
	"time"
)

func (repo *URLAliasRepo) GetAllExpiredAliasIDs(ctx context.Context) ([]uint, error) {
	acc := []uint{}
	rn := time.Now().Unix()

	rows, err := repo.db.QueryContext(ctx, "SELECT ID FROM URLAliases WHERE DeleteAt <= ?", rn)
	if err != nil {
		return acc, err
	}

	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return acc, err
		}

		acc = append(acc, id)
	}

	return acc, rows.Err()
} 
