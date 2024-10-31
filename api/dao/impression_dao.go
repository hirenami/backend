package dao

import (
	"context"
	"database/sql"
)

func (d *Dao) PlusImpression(ctx context.Context, tx *sql.Tx, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.PlusImpression(ctx, tweetid)
}

func (d *Dao) GetImpression(ctx context.Context, tx *sql.Tx, tweetid int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetImpression(ctx, tweetid)
}