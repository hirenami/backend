package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) GetFollowRequest(ctx context.Context, tx *sql.Tx, followerid string) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetFollowRequest(ctx, followerid)
}

func (d *Dao) CreateKeyFollow(ctx context.Context, tx *sql.Tx, followerid, followingid string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateKeyFollowParams{
		Followerid:  followerid,
		Followingid: followingid,
	}

	return txQueries.CreateKeyFollow(ctx, arg)
}

func (d *Dao) DeleteKeyFollow(ctx context.Context, tx *sql.Tx, followerid, followingid string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.DeleteKeyFollowParams{
		Followerid:  followerid,
		Followingid: followingid,
	}

	return txQueries.DeleteKeyFollow(ctx, arg)
}
