package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CheckLike(ctx context.Context, tx *sql.Tx, tweetid int32) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.CheckLike(ctx, tweetid)
}

func (d *Dao) CreateLike(ctx context.Context, tx *sql.Tx, userid string, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateLikeParams{
		Userid:  userid,
		Tweetid: tweetid,
	}

	return txQueries.CreateLike(ctx, args)
}

func (d *Dao) DeleteLike(ctx context.Context, tx *sql.Tx, userid string, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.DeleteLikeParams{
		Userid:  userid,
		Tweetid: tweetid,
	}

	return txQueries.DeleteLike(ctx, args)
}

func (d *Dao) GetLikes(ctx context.Context, tx *sql.Tx, tweetid int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetLikes(ctx, tweetid)
}

func (d *Dao) MinusLike(ctx context.Context, tx *sql.Tx, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.MinusLike(ctx, tweetid)
}

func (d *Dao) PlusLike(ctx context.Context, tx *sql.Tx, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.PlusLike(ctx, tweetid)
}

func (d *Dao) GetLikesUser(ctx context.Context, tx *sql.Tx, userId string) ([]int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetUsersLikes(ctx, userId)
}

func (d *Dao) IsLiked(ctx context.Context, tx *sql.Tx, userId string, tweetId int32) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)
	arg := sqlc.IsLikedParams{
		Userid:  userId,
		Tweetid: tweetId,
	}

	return txQueries.IsLiked(ctx, arg)
}
