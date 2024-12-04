// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: setting.sql

package sqlc

import (
	"context"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO users (firebaseUid ,userId, username) VALUES (?, ?, ?)
`

type CreateAccountParams struct {
	Firebaseuid string `json:"firebaseuid"`
	Userid      string `json:"userid"`
	Username    string `json:"username"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount, arg.Firebaseuid, arg.Userid, arg.Username)
	return err
}

const createIsAdmin = `-- name: CreateIsAdmin :exec
UPDATE users
SET isAdmin = ?
WHERE userId = ?
`

type CreateIsAdminParams struct {
	Isadmin bool   `json:"isadmin"`
	Userid  string `json:"userid"`
}

func (q *Queries) CreateIsAdmin(ctx context.Context, arg CreateIsAdminParams) error {
	_, err := q.db.ExecContext(ctx, createIsAdmin, arg.Isadmin, arg.Userid)
	return err
}

const createIsPremium = `-- name: CreateIsPremium :exec
UPDATE users
SET isPremium = true
WHERE userId = ?
`

func (q *Queries) CreateIsPremium(ctx context.Context, userid string) error {
	_, err := q.db.ExecContext(ctx, createIsPremium, userid)
	return err
}

const createIsPrivate = `-- name: CreateIsPrivate :exec
UPDATE users
SET isPrivate = ?
WHERE userId = ?
`

type CreateIsPrivateParams struct {
	Isprivate bool   `json:"isprivate"`
	Userid    string `json:"userid"`
}

func (q *Queries) CreateIsPrivate(ctx context.Context, arg CreateIsPrivateParams) error {
	_, err := q.db.ExecContext(ctx, createIsPrivate, arg.Isprivate, arg.Userid)
	return err
}

const getIdByUID = `-- name: GetIdByUID :one
SELECT userId FROM users WHERE firebaseUid = ?
`

func (q *Queries) GetIdByUID(ctx context.Context, firebaseuid string) (string, error) {
	row := q.db.QueryRowContext(ctx, getIdByUID, firebaseuid)
	var userid string
	err := row.Scan(&userid)
	return userid, err
}

const isUserExists = `-- name: IsUserExists :one
SELECT EXISTS (
    SELECT 1 
    FROM users 
    WHERE userId = ?
)
`

func (q *Queries) IsUserExists(ctx context.Context, userid string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isUserExists, userid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateUsername = `-- name: UpdateUsername :exec
UPDATE users
SET username = ?
WHERE userId = ?
`

type UpdateUsernameParams struct {
	Username string `json:"username"`
	Userid   string `json:"userid"`
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) error {
	_, err := q.db.ExecContext(ctx, updateUsername, arg.Username, arg.Userid)
	return err
}
