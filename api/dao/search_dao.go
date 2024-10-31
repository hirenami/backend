package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) SearchByHashtag(ctx context.Context, tx * sql.Tx, hashtag string) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.SearchByHashtag(ctx, hashtag)
}

func (d *Dao) SearchByKeyword(ctx context.Context, tx * sql.Tx, concat interface{}) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.SearchByKeyword(ctx, concat)
}

func (d *Dao) SearchUser(ctx context.Context, tx * sql.Tx, concat interface{}) ([]sqlc.User, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.SearchUser(ctx, concat)
}

