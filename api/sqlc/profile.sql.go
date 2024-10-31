// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: profile.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createBiography = `-- name: CreateBiography :exec
UPDATE users
SET biography = ?
WHERE userId = ?
`

type CreateBiographyParams struct {
	Biography sql.NullString `json:"biography"`
	Userid    string         `json:"userid"`
}

func (q *Queries) CreateBiography(ctx context.Context, arg CreateBiographyParams) error {
	_, err := q.db.ExecContext(ctx, createBiography, arg.Biography, arg.Userid)
	return err
}

const createHeaderImage = `-- name: CreateHeaderImage :exec
UPDATE users
SET header_image = ?
WHERE userId = ?
`

type CreateHeaderImageParams struct {
	HeaderImage string `json:"header_image"`
	Userid      string `json:"userid"`
}

func (q *Queries) CreateHeaderImage(ctx context.Context, arg CreateHeaderImageParams) error {
	_, err := q.db.ExecContext(ctx, createHeaderImage, arg.HeaderImage, arg.Userid)
	return err
}

const createIconImage = `-- name: CreateIconImage :exec
UPDATE users
SET icon_image = ?
WHERE userId = ?
`

type CreateIconImageParams struct {
	IconImage string `json:"icon_image"`
	Userid    string `json:"userid"`
}

func (q *Queries) CreateIconImage(ctx context.Context, arg CreateIconImageParams) error {
	_, err := q.db.ExecContext(ctx, createIconImage, arg.IconImage, arg.Userid)
	return err
}

const getProfile = `-- name: GetProfile :one
SELECT firebaseuid, userid, username, created_at, header_image, icon_image, biography, isprivate, isfrozen, isdeleted, isadmin FROM users WHERE userId = ?
`

func (q *Queries) GetProfile(ctx context.Context, userid string) (User, error) {
	row := q.db.QueryRowContext(ctx, getProfile, userid)
	var i User
	err := row.Scan(
		&i.Firebaseuid,
		&i.Userid,
		&i.Username,
		&i.CreatedAt,
		&i.HeaderImage,
		&i.IconImage,
		&i.Biography,
		&i.Isprivate,
		&i.Isfrozen,
		&i.Isdeleted,
		&i.Isadmin,
	)
	return i, err
}
