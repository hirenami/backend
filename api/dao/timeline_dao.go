package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) Timeline(ctx context.Context, tx *sql.Tx, followingid string) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.TimelineParams{
		Followingid: followingid,
		Followingid_2: followingid,
	}

	return txQueries.Timeline(ctx, args)
}