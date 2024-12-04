package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateAccount(ctx context.Context, tx *sql.Tx, firebaseUid string, username string, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateAccountParams{
		Username: username,
		Userid:   userId,
		Firebaseuid:      firebaseUid,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateAccount(ctx, arg)
}

func (d *Dao) CreateIsAdmin(ctx context.Context, tx *sql.Tx, Isadmin bool, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateIsAdminParams{
		Isadmin: Isadmin,
		Userid:  userId,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateIsAdmin(ctx, arg)
}

func (d *Dao) CreateIsPrivate(ctx context.Context, tx *sql.Tx, Isprivate bool, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateIsPrivateParams{
		Isprivate: Isprivate,
		Userid:    userId,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateIsPrivate(ctx, arg)
}

func (d *Dao) CreateIsPremium(ctx context.Context, tx *sql.Tx,  userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.CreateIsPremium(ctx, userId)
}

func (d *Dao) IsUserExists(ctx context.Context, tx *sql.Tx, userId string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.IsUserExists(ctx, userId)
}

func (d *Dao) UpdateUsername(ctx context.Context, tx *sql.Tx, username string, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.UpdateUsernameParams{
		Username: username,
		Userid:   userId,
	}

	// トランザクション内でクエリを実行
	return txQueries.UpdateUsername(ctx, arg)
}

func (d *Dao) GetIdbyUid(ctx context.Context, tx *sql.Tx, firebaseUid string) (string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetIdByUID(ctx, firebaseUid)
}
