package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateTweet(ctx context.Context, tx *sql.Tx, userId, content, media_url string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateTweetParams{
		Userid:   userId,
		Content:  content,
		MediaUrl: media_url,
	}

	return txQueries.CreateTweet(ctx, args)
}

func (d *Dao) CreateRetweet(ctx context.Context, tx *sql.Tx, userId string, retweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateRetweetParams{
		Userid: userId,
		Retweetid: retweetId,
	}

	return txQueries.CreateRetweet(ctx, args)
}

func (d *Dao) CreateQuote(ctx context.Context, tx *sql.Tx, userId, content, media_url string, retweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateQuoteParams{
		Userid:   userId,
		Content:  content,
		MediaUrl: media_url,
		Retweetid: retweetId,
	}

	return txQueries.CreateQuote(ctx, args)
}

func (d *Dao) DeleteTweet(ctx context.Context, tx *sql.Tx, tweetid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.DeleteTweet(ctx, tweetid)
}

func (d *Dao) EditTweet(ctx context.Context, tx *sql.Tx, content, media_url string, tweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.EditTweetParams{
		Content:  content,
		MediaUrl: media_url,
		Tweetid:  tweetId,
	}

	return txQueries.EditTweet(ctx, args)
}

func (d *Dao) GetQuotes(ctx context.Context, tx *sql.Tx, retweetId int32) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetQuotes(ctx, retweetId)

}

func (d *Dao) GetRetweets(ctx context.Context, tx *sql.Tx, retweetId int32) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)


	return txQueries.GetRetweets(ctx, retweetId)
}

func (d *Dao) GetRetweetsCount(ctx context.Context, tx *sql.Tx, tweetId int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetRetweetsCount(ctx, tweetId)
}

func (d *Dao) GetUsersTweet(ctx context.Context, tx *sql.Tx, userId string) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetUsersTweets(ctx, userId)
}

func (d *Dao) IsTweetExists(ctx context.Context, tx *sql.Tx, tweetId int32) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.IsTweetExists(ctx, tweetId)
}

func (d *Dao) MinusRetweet(ctx context.Context, tx *sql.Tx, tweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.MinusRetweet(ctx, tweetId)
}

func (d *Dao) PlusRetweet(ctx context.Context, tx *sql.Tx, tweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.PlusRetweet(ctx, tweetId)
}

func (d *Dao) GetUserId(ctx context.Context, tx *sql.Tx, tweetId int32) (string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetUserId(ctx, tweetId)
}

func (d *Dao) GetRetweetId(ctx context.Context, tx *sql.Tx, tweetId int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetRetweetId(ctx, tweetId)
}

func (d *Dao) GetTweet(ctx context.Context, tx *sql.Tx, tweetId int32) (sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetTweet(ctx, tweetId)
}

func (d *Dao) IsRetweet (ctx context.Context, tx *sql.Tx,userId string, retweetId int32) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.IsRetweetParams{
		Userid: userId,
		Retweetid: retweetId,
	}

	return txQueries.IsRetweet(ctx, arg)
}

func (d *Dao) GetTweetId(ctx context.Context, tx *sql.Tx, userId string, retweetId int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.GetTweetIdParams{
		Userid: userId,
		Retweetid: retweetId,
	}

	return txQueries.GetTweetId(ctx, arg)
}
