// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: timeline.sql

package sqlc

import (
	"context"
)

const timeline = `-- name: Timeline :many
SELECT tweetid, userid, retweetid, isquote, isreply, review, created_at, updated_at, content, media_url, likes, retweets, replies, impressions, isdeleted FROM tweets
WHERE userId IN (
	SELECT followerId FROM follows
	WHERE followingId = ?
) AND isDeleted = false AND isReply = false ORDER BY created_at DESC
`

func (q *Queries) Timeline(ctx context.Context, followingid string) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, timeline, followingid)
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
