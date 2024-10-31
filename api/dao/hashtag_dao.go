package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateHashtag(ctx context.Context, tx *sql.Tx, hashtag string, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateHashtagParams{
		Hashtag: hashtag,
		Tweetid: tweetid,
	}
	return txQueries.CreateHashtag(ctx, arg)
}

func (d *Dao) UpdateHashtag(ctx context.Context, tx *sql.Tx, hashtag string, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.UpdateHashtagParams{
		Hashtag: hashtag,
		Tweetid: tweetid,
	}
	return txQueries.UpdateHashtag(ctx, arg)
}

func (d *Dao) DeleteHashtag(ctx context.Context, tx *sql.Tx, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.DeleteHashtag(ctx, tweetid)
}