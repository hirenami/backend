package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) Timeline(ctx context.Context, tx *sql.Tx, followingid string) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.Timeline(ctx, followingid)
}

// おすすめツイートも実装予定