// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tweets.sql

package sqlc

import (
	"context"
)

const createQuote = `-- name: CreateQuote :exec
INSERT INTO tweets (
	userId, retweetId, content, media_url, isQuote
) VALUES (
	?, ?, ?, ?, true
)
`

type CreateQuoteParams struct {
	Userid    string `json:"userid"`
	Retweetid int32  `json:"retweetid"`
	Content   string `json:"content"`
	MediaUrl  string `json:"media_url"`
}

func (q *Queries) CreateQuote(ctx context.Context, arg CreateQuoteParams) error {
	_, err := q.db.ExecContext(ctx, createQuote,
		arg.Userid,
		arg.Retweetid,
		arg.Content,
		arg.MediaUrl,
	)
	return err
}

const createRetweet = `-- name: CreateRetweet :exec
INSERT INTO tweets (
	userId, retweetId, content
) VALUES (
	?, ? , ''
)
`

type CreateRetweetParams struct {
	Userid    string `json:"userid"`
	Retweetid int32  `json:"retweetid"`
}

func (q *Queries) CreateRetweet(ctx context.Context, arg CreateRetweetParams) error {
	_, err := q.db.ExecContext(ctx, createRetweet, arg.Userid, arg.Retweetid)
	return err
}

const createTweet = `-- name: CreateTweet :exec
INSERT INTO tweets (
	userId, content, media_url
) VALUES (
	?, ?, ?
)
`

type CreateTweetParams struct {
	Userid   string `json:"userid"`
	Content  string `json:"content"`
	MediaUrl string `json:"media_url"`
}

func (q *Queries) CreateTweet(ctx context.Context, arg CreateTweetParams) error {
	_, err := q.db.ExecContext(ctx, createTweet, arg.Userid, arg.Content, arg.MediaUrl)
	return err
}

const deleteTweet = `-- name: DeleteTweet :exec
UPDATE tweets
SET isDeleted = true
WHERE tweetId = ?
`

func (q *Queries) DeleteTweet(ctx context.Context, tweetid int32) error {
	_, err := q.db.ExecContext(ctx, deleteTweet, tweetid)
	return err
}

const editTweet = `-- name: EditTweet :exec
UPDATE tweets
SET content = ?, media_url = ?
WHERE tweetId = ?
`

type EditTweetParams struct {
	Content  string `json:"content"`
	MediaUrl string `json:"media_url"`
	Tweetid  int32  `json:"tweetid"`
}

func (q *Queries) EditTweet(ctx context.Context, arg EditTweetParams) error {
	_, err := q.db.ExecContext(ctx, editTweet, arg.Content, arg.MediaUrl, arg.Tweetid)
	return err
}

const getQuotes = `-- name: GetQuotes :many
SELECT tweetid, userid, retweetid, isquote, isreply, review, created_at, updated_at, content, media_url, likes, retweets, replies, impressions, isdeleted FROM tweets
WHERE retweetId = ? and isQuote = true and isDeleted = false ORDER BY created_at DESC
`

func (q *Queries) GetQuotes(ctx context.Context, retweetid int32) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getQuotes, retweetid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tweet{}
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.Tweetid,
			&i.Userid,
			&i.Retweetid,
			&i.Isquote,
			&i.Isreply,
			&i.Review,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Content,
			&i.MediaUrl,
			&i.Likes,
			&i.Retweets,
			&i.Replies,
			&i.Impressions,
			&i.Isdeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRetweetId = `-- name: GetRetweetId :one
SELECT retweetId FROM tweets
WHERE tweetId = ?
`

func (q *Queries) GetRetweetId(ctx context.Context, tweetid int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRetweetId, tweetid)
	var retweetid int32
	err := row.Scan(&retweetid)
	return retweetid, err
}

const getRetweets = `-- name: GetRetweets :many
SELECT userId FROM tweets
WHERE retweetId = ? and isQuote = false and isDeleted = false ORDER BY created_at DESC
`

func (q *Queries) GetRetweets(ctx context.Context, retweetid int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getRetweets, retweetid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var userid string
		if err := rows.Scan(&userid); err != nil {
			return nil, err
		}
		items = append(items, userid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRetweetsCount = `-- name: GetRetweetsCount :one
SELECT retweets FROM tweets
WHERE tweetId = ? and isDeleted=false
`

func (q *Queries) GetRetweetsCount(ctx context.Context, tweetid int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRetweetsCount, tweetid)
	var retweets int32
	err := row.Scan(&retweets)
	return retweets, err
}

const getTweet = `-- name: GetTweet :one
SELECT tweetid, userid, retweetid, isquote, isreply, review, created_at, updated_at, content, media_url, likes, retweets, replies, impressions, isdeleted FROM tweets
WHERE tweetId = ?
`

func (q *Queries) GetTweet(ctx context.Context, tweetid int32) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, getTweet, tweetid)
	var i Tweet
	err := row.Scan(
		&i.Tweetid,
		&i.Userid,
		&i.Retweetid,
		&i.Isquote,
		&i.Isreply,
		&i.Review,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.MediaUrl,
		&i.Likes,
		&i.Retweets,
		&i.Replies,
		&i.Impressions,
		&i.Isdeleted,
	)
	return i, err
}

const getTweetId = `-- name: GetTweetId :one
SELECT tweetId FROM tweets
WHERE retweetId = ? and userId = ? and isDeleted = false and isQuote = false
`

type GetTweetIdParams struct {
	Retweetid int32  `json:"retweetid"`
	Userid    string `json:"userid"`
}

func (q *Queries) GetTweetId(ctx context.Context, arg GetTweetIdParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getTweetId, arg.Retweetid, arg.Userid)
	var tweetid int32
	err := row.Scan(&tweetid)
	return tweetid, err
}

const getUserId = `-- name: GetUserId :one
SELECT userId FROM tweets
WHERE tweetId = ?
`

func (q *Queries) GetUserId(ctx context.Context, tweetid int32) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserId, tweetid)
	var userid string
	err := row.Scan(&userid)
	return userid, err
}

const getUsersTweets = `-- name: GetUsersTweets :many
SELECT tweetid, userid, retweetid, isquote, isreply, review, created_at, updated_at, content, media_url, likes, retweets, replies, impressions, isdeleted FROM tweets 
WHERE userId = ? and isReply = false and isDeleted = false ORDER BY created_at DESC
`

func (q *Queries) GetUsersTweets(ctx context.Context, userid string) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getUsersTweets, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tweet{}
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.Tweetid,
			&i.Userid,
			&i.Retweetid,
			&i.Isquote,
			&i.Isreply,
			&i.Review,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Content,
			&i.MediaUrl,
			&i.Likes,
			&i.Retweets,
			&i.Replies,
			&i.Impressions,
			&i.Isdeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isRetweet = `-- name: IsRetweet :one
SELECT EXISTS (
	SELECT 1 
	FROM tweets 
	WHERE retweetId = ? and isDeleted = false and userId = ? and isQuote = false
)
`

type IsRetweetParams struct {
	Retweetid int32  `json:"retweetid"`
	Userid    string `json:"userid"`
}

func (q *Queries) IsRetweet(ctx context.Context, arg IsRetweetParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isRetweet, arg.Retweetid, arg.Userid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isTweetExists = `-- name: IsTweetExists :one
SELECT EXISTS (
    SELECT 1 
    FROM tweets 
	WHERE tweetId = ?
)
`

func (q *Queries) IsTweetExists(ctx context.Context, tweetid int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, isTweetExists, tweetid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const minusRetweet = `-- name: MinusRetweet :exec
UPDATE tweets
SET retweets = retweets - 1
WHERE tweetId = ?
`

func (q *Queries) MinusRetweet(ctx context.Context, tweetid int32) error {
	_, err := q.db.ExecContext(ctx, minusRetweet, tweetid)
	return err
}

const plusRetweet = `-- name: PlusRetweet :exec
UPDATE tweets
SET retweets = retweets + 1
WHERE tweetId = ?
`

func (q *Queries) PlusRetweet(ctx context.Context, tweetid int32) error {
	_, err := q.db.ExecContext(ctx, plusRetweet, tweetid)
	return err
}

const updateReview = `-- name: UpdateReview :exec
UPDATE tweets
SET review = ?
WHERE tweetId = ?
`

type UpdateReviewParams struct {
	Review  int32 `json:"review"`
	Tweetid int32 `json:"tweetid"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) error {
	_, err := q.db.ExecContext(ctx, updateReview, arg.Review, arg.Tweetid)
	return err
}
