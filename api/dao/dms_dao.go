package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateDm(ctx context.Context, tx *sql.Tx, senderId, receiverId, content, media_url string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	var mediaUrl sql.NullString
	if media_url == "" {
		mediaUrl = sql.NullString{
			String: "",
			Valid:  false, // NULLを示す
		}
	} else {
		mediaUrl = sql.NullString{
			String: media_url,
			Valid:  true, // 有効なURL
		}
	}

	arg := sqlc.CreateDmParams{
		Senderid:   senderId,
		Receiverid: receiverId,
		Content:    content,
		MediaUrl:   mediaUrl,
	}
	return txQueries.CreateDm(ctx, arg)
}

func (d *Dao) DeleteDm(ctx context.Context, tx *sql.Tx, dmsId int32) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.DeleteDm(ctx, dmsId)
}

func (d *Dao) GetDm(ctx context.Context, tx *sql.Tx, dmsid int32) (sqlc.Dm, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetDm(ctx, dmsid)
}

func (d *Dao) GetDms(ctx context.Context, tx *sql.Tx, senderId, receiverId string) ([]sqlc.Dm, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.GetDmsParams{
		Senderid:     senderId,
		Receiverid:   receiverId,
		Senderid_2:   receiverId,
		Receiverid_2: senderId,
	}

	return txQueries.GetDms(ctx, arg)
}

func (d *Dao) GetLastMessages(ctx context.Context, tx *sql.Tx, senderId, receiverId string) ([]sqlc.Dm, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.GetLastMessagesParams{
		Senderid:     senderId,
		Receiverid:   receiverId,
		Senderid_2:   receiverId,
		Receiverid_2: senderId,
	}

	return txQueries.GetLastMessages(ctx, arg)
}

func (d *Dao) GetDmsUsers(ctx context.Context, tx *sql.Tx, userId string) ([]interface{}, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.GetDmsUsersParams{
		Senderid:   userId,
		Senderid_2: userId,
		Receiverid: userId,
	}
	return txQueries.GetDmsUsers(ctx, args)
}

func (d *Dao) SetDmStatus(ctx context.Context, tx *sql.Tx, senderId, receiverId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	args := sqlc.SetDmStatusParams{
		Senderid:     senderId,
		Senderid_2:   receiverId,
		Receiverid:   receiverId,
		Receiverid_2: senderId,
	}

	return txQueries.SetDmStatus(ctx, args)
}
