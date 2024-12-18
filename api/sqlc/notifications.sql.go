// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: notifications.sql

package sqlc

import (
	"context"
)

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (
	senderId, replyId, type, contentId
) VALUES (
	?, ?, ?, ?
)
`

type CreateNotificationParams struct {
	Senderid  string `json:"senderid"`
	Replyid   string `json:"replyid"`
	Type      string `json:"type"`
	Contentid int32  `json:"contentid"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) error {
	_, err := q.db.ExecContext(ctx, createNotification,
		arg.Senderid,
		arg.Replyid,
		arg.Type,
		arg.Contentid,
	)
	return err
}

const getNotification = `-- name: GetNotification :one
SELECT notificationsid, senderid, replyid, type, createdat, status, contentid FROM notifications WHERE notificationsId = ?
`

func (q *Queries) GetNotification(ctx context.Context, notificationsid int32) (Notification, error) {
	row := q.db.QueryRowContext(ctx, getNotification, notificationsid)
	var i Notification
	err := row.Scan(
		&i.Notificationsid,
		&i.Senderid,
		&i.Replyid,
		&i.Type,
		&i.Createdat,
		&i.Status,
		&i.Contentid,
	)
	return i, err
}

const getNotifications = `-- name: GetNotifications :many
SELECT notificationsid, senderid, replyid, type, createdat, status, contentid FROM notifications WHERE replyId = ? ORDER BY createdAt DESC
`

func (q *Queries) GetNotifications(ctx context.Context, replyid string) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getNotifications, replyid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Notification{}
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.Notificationsid,
			&i.Senderid,
			&i.Replyid,
			&i.Type,
			&i.Createdat,
			&i.Status,
			&i.Contentid,
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

const isNotificationExist = `-- name: IsNotificationExist :one
SELECT EXISTS (
    SELECT 1 
    FROM notifications 
	WHERE notificationsId = ?
)
`

func (q *Queries) IsNotificationExist(ctx context.Context, notificationsid int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, isNotificationExist, notificationsid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateNotificationStatus = `-- name: UpdateNotificationStatus :exec
UPDATE notifications
SET status = 'read'
WHERE notificationsId = ?
`

func (q *Queries) UpdateNotificationStatus(ctx context.Context, notificationsid int32) error {
	_, err := q.db.ExecContext(ctx, updateNotificationStatus, notificationsid)
	return err
}
