package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateFollow(ctx context.Context, tx *sql.Tx, followerId, followingId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateFollowParams{
		Followerid:  followerId,
		Followingid: followingId,
	}
	return txQueries.CreateFollow(ctx, arg)
}

func (d *Dao) CountFollower(ctx context.Context, tx *sql.Tx, followerId string) (int64, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.CountFollower(ctx, followerId)
}

func (d *Dao) CountFollowing(ctx context.Context, tx *sql.Tx, followingId string) (int64, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.CountFollowing(ctx, followingId)
}

func (d *Dao) GetFollower(ctx context.Context, tx *sql.Tx, followerId string) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetFollower(ctx, followerId)
}

func (d *Dao) GetFollowing(ctx context.Context, tx *sql.Tx, followingId string) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetFollowing(ctx, followingId)
}

func (d *Dao) DeleteFollow(ctx context.Context, tx *sql.Tx, followerId, followingId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.DeleteFollowParams{
		Followerid:  followerId,
		Followingid: followingId,
	}
	return txQueries.DeleteFollow(ctx, arg)
}

func (d *Dao) IsFollowing(ctx context.Context, tx *sql.Tx, followerId, followingId string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.IsFollowingParams{
		Followerid:  followerId,
		Followingid: followingId,
	}

	return txQueries.IsFollowing(ctx, arg)
}