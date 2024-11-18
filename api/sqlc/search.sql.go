// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: search.sql

package sqlc

import (
	"context"
)

const searchByHashtag = `-- name: SearchByHashtag :many
SELECT tweetid, userid, retweetid, isquote, isreply, isreview, created_at, updated_at, content, media_url, likes, retweets, replies, impressions, isdeleted FROM tweets
WHERE tweetId IN (
	SELECT tweetId FROM hashtags
	WHERE hashtag = ?
) AND isDeleted = false ORDER BY created_at DESC
`

func (q *Queries) SearchByHashtag(ctx context.Context, hashtag string) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, searchByHashtag, hashtag)
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
			&i.Isreview,
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

const searchByKeyword = `-- name: SearchByKeyword :many
SELECT tweetid, userid, retweetid, isquote, isreply, isreview, created_at, updated_at, content, media_url, likes, retweets, replies, impressions, isdeleted FROM tweets
WHERE content LIKE CONCAT('%', ? , '%') AND isDeleted = false ORDER BY created_at DESC
`

func (q *Queries) SearchByKeyword(ctx context.Context, concat interface{}) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, searchByKeyword, concat)
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
			&i.Isreview,
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

const searchUser = `-- name: SearchUser :many
SELECT firebaseuid, userid, username, created_at, header_image, icon_image, biography, isprivate, ispremium, isdeleted, isadmin FROM users
WHERE username LIKE CONCAT('%', ? , '%') ORDER BY created_at DESC
`

func (q *Queries) SearchUser(ctx context.Context, concat interface{}) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, searchUser, concat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Firebaseuid,
			&i.Userid,
			&i.Username,
			&i.CreatedAt,
			&i.HeaderImage,
			&i.IconImage,
			&i.Biography,
			&i.Isprivate,
			&i.Ispremium,
			&i.Isdeleted,
			&i.Isadmin,
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
