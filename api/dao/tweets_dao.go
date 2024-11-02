package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateTweet(ctx context.Context, tx *sql.Tx, userId, content, media_url string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	var mediaUrl sql.NullString
	if media_url == "" {
		mediaUrl = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		mediaUrl = sql.NullString{
			String: media_url,
			Valid:  true,
		}
	}

	args := sqlc.CreateTweetParams{
		Userid:   userId,
		Content:  content,
		MediaUrl: mediaUrl,
	}

	return txQueries.CreateTweet(ctx, args)
}

func (d *Dao) CreateRetweet(ctx context.Context, tx *sql.Tx, userId string, retweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateRetweetParams{
		Userid: userId,
		Retweetid: sql.NullInt32{
			Int32: retweetId,
			Valid: true,
		},
	}

	return txQueries.CreateRetweet(ctx, args)
}

func (d *Dao) CreateQuote(ctx context.Context, tx *sql.Tx, userId, content, media_url string, retweetId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	var mediaUrl sql.NullString
	if media_url == "" {
		mediaUrl = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		mediaUrl = sql.NullString{
			String: media_url,
			Valid:  true,
		}
	}

	args := sqlc.CreateQuoteParams{
		Userid:   userId,
		Content:  content,
		MediaUrl: mediaUrl,
		Retweetid: sql.NullInt32{
			Int32: retweetId,
			Valid: true,
		},
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

	var mediaUrl sql.NullString
	if media_url == "" {
		mediaUrl = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		mediaUrl = sql.NullString{
			String: media_url,
			Valid:  true,
		}
	}

	args := sqlc.EditTweetParams{
		Content:  content,
		MediaUrl: mediaUrl,
		Tweetid:  tweetId,
	}

	return txQueries.EditTweet(ctx, args)
}

func (d *Dao) GetQuotes(ctx context.Context, tx *sql.Tx, retweetId int32) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	var RetweetId sql.NullInt32
	if retweetId == 0 {
		RetweetId = sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
	} else {
		RetweetId = sql.NullInt32{
			Int32: retweetId,
			Valid: true,
		}
	}

	return txQueries.GetQuotes(ctx, RetweetId)

}

func (d *Dao) GetRetweets(ctx context.Context, tx *sql.Tx, retweetId int32) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	var RetweetId sql.NullInt32
	if retweetId == 0 {
		RetweetId = sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
	} else {
		RetweetId = sql.NullInt32{
			Int32: retweetId,
			Valid: true,
		}
	}

	return txQueries.GetRetweets(ctx, RetweetId)
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

func (d *Dao) GetRetweetId(ctx context.Context, tx *sql.Tx, tweetId int32) (sql.NullInt32, error) {
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

	RetweetId := sql.NullInt32{
		Int32: retweetId,
		Valid: true,
	}

	arg := sqlc.IsRetweetParams{
		Userid: userId,
		Retweetid: RetweetId,
	}

	return txQueries.IsRetweet(ctx, arg)
}

func (d *Dao) GetTweetId(ctx context.Context, tx *sql.Tx, userId string, retweetId int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	RetweetId := sql.NullInt32{
		Int32: retweetId,
		Valid: true,
	}

	arg := sqlc.GetTweetIdParams{
		Userid: userId,
		Retweetid: RetweetId,
	}

	return txQueries.GetTweetId(ctx, arg)
}
