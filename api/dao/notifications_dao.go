package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateNotification(ctx context.Context,tx *sql.Tx, senderId, replyId, Type string, contentId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.CreateNotificationParams{
		Senderid: senderId,
		Replyid:  replyId,
		Type:     Type,
		Contentid: contentId,
	}

	return txQueries.CreateNotification(ctx, args)
}

func (d *Dao) GetNotification(ctx context.Context, tx *sql.Tx,notificationsid int32) (sqlc.Notification, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetNotification(ctx, notificationsid)
}

func (d *Dao) GetNotifications(ctx context.Context, tx *sql.Tx, repid string) ([]sqlc.Notification, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetNotifications(ctx, repid)
}

func (d *Dao) UpdateNotificationStatus(ctx context.Context, tx *sql.Tx, notificationsid int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.UpdateNotificationStatus(ctx, notificationsid)
}

func (d *Dao) IsNotificationExists(ctx context.Context, tx *sql.Tx, notificationsid int32) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.IsNotificationExist(ctx, notificationsid)
}