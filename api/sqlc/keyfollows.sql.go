// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: keyfollows.sql

package sqlc

import (
	"context"
)

const createKeyFollow = `-- name: CreateKeyFollow :exec
INSERT keyfollows (
	followerId, followingId
) VALUES (
	?, ?
)
`

type CreateKeyFollowParams struct {
	Followerid  string `json:"followerid"`
	Followingid string `json:"followingid"`
}

func (q *Queries) CreateKeyFollow(ctx context.Context, arg CreateKeyFollowParams) error {
	_, err := q.db.ExecContext(ctx, createKeyFollow, arg.Followerid, arg.Followingid)
	return err
}

const deleteKeyFollow = `-- name: DeleteKeyFollow :exec
DELETE FROM keyfollows
WHERE followerId = ? AND followingId = ?
`

type DeleteKeyFollowParams struct {
	Followerid  string `json:"followerid"`
	Followingid string `json:"followingid"`
}

func (q *Queries) DeleteKeyFollow(ctx context.Context, arg DeleteKeyFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteKeyFollow, arg.Followerid, arg.Followingid)
	return err
}

const deleteKeyFollows = `-- name: DeleteKeyFollows :exec
DELETE FROM keyfollows
WHERE followerId = ?
`

func (q *Queries) DeleteKeyFollows(ctx context.Context, followerid string) error {
	_, err := q.db.ExecContext(ctx, deleteKeyFollows, followerid)
	return err
}

const getFollowRequest = `-- name: GetFollowRequest :many
SELECT followingId FROM keyfollows WHERE followerId = ?
`

func (q *Queries) GetFollowRequest(ctx context.Context, followerid string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowRequest, followerid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var followingid string
		if err := rows.Scan(&followingid); err != nil {
			return nil, err
		}
		items = append(items, followingid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isFollowRequest = `-- name: IsFollowRequest :one
SELECT EXISTS (
	SELECT 1
	FROM keyfollows
	WHERE followerId = ? AND followingId = ?
)
`

type IsFollowRequestParams struct {
	Followerid  string `json:"followerid"`
	Followingid string `json:"followingid"`
}

func (q *Queries) IsFollowRequest(ctx context.Context, arg IsFollowRequestParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isFollowRequest, arg.Followerid, arg.Followingid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
