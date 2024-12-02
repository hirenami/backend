package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateReply(ctx context.Context, tx *sql.Tx, userId, content, mediaUrl string, review int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateReplyParams{
		Userid:   userId,
		Content:  content,
		MediaUrl: mediaUrl,
		Review:   review,
	}
	return txQueries.CreateReply(ctx, args)
}

func (d *Dao) GetRepliesToTweet(ctx context.Context, tx *sql.Tx, tweetID int32) ([]int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetRepliesToTweet(ctx, tweetID)
}

func (d *Dao) GetTweetRepliedTo(ctx context.Context, tx *sql.Tx, replyID int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetTweetRepliedTo(ctx, replyID)
}

func (d *Dao) GetUsersReplies(ctx context.Context, tx *sql.Tx, userId string) ([]sqlc.Tweet, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetUsersReplies(ctx, userId)
}

func (d *Dao) RelateReplyToTweet(ctx context.Context, tx *sql.Tx, tweetID, replyID int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.RelateReplyToTweetParams{
		Tweetid: tweetID,
		Replyid: replyID,
	}
	return txQueries.RelateReplyToTweet(ctx, args)
}

func (d *Dao) DeleteRelateReplyToTweet(ctx context.Context, tx *sql.Tx, tweetID, replyID int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.UnrelateReplyToTweetParams{
		Tweetid: tweetID,
		Replyid: replyID,
	}
	return txQueries.UnrelateReplyToTweet(ctx, args)
}

func (d *Dao) GetLastInsertID(ctx context.Context, tx *sql.Tx) (int64, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetLastInsertID(ctx)
}

func (d *Dao) PlusReplyCount(ctx context.Context, tx *sql.Tx, tweetID int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.PlusOneReply(ctx, tweetID)
}

func (d *Dao) MinusReplyCount(ctx context.Context, tx *sql.Tx, tweetID int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.MinusOneReply(ctx, tweetID)
}

func (d *Dao) CountReplies(ctx context.Context, tx *sql.Tx, replyID int32) (int32, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.CountReplies(ctx, replyID)
}

func (d *Dao) IsReplyExists(ctx context.Context, tx *sql.Tx, replyID int32) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.IsReplyExists(ctx, replyID)
}
