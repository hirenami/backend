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

func (d *Dao) CreateIsDeleted(ctx context.Context, tx *sql.Tx, Isadmin bool, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateIsDeletedParams{
		Isdeleted: Isadmin,
		Userid:    userId,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateIsDeleted(ctx, arg)
}

func (d *Dao) CreateIsFrozen(ctx context.Context, tx *sql.Tx, Isadmin bool, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateIsFrozenParams{
		Isfrozen: Isadmin,
		Userid:   userId,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateIsFrozen(ctx, arg)
}

func (d *Dao) CreateIsPrivate(ctx context.Context, tx *sql.Tx, Isadmin bool, userId string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateIsPrivateParams{
		Isprivate: Isadmin,
		Userid:    userId,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateIsPrivate(ctx, arg)
}

func (d *Dao) GetIsAdmin(ctx context.Context, tx *sql.Tx, userId string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetIsAdmin(ctx, userId)
}

func (d *Dao) GetIsDeleted(ctx context.Context, tx *sql.Tx, userId string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetIsDeleted(ctx, userId)
}

func (d *Dao) GetIsFrozen(ctx context.Context, tx *sql.Tx, userId string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetIsFrozen(ctx, userId)
}

func (d *Dao) GetIsPrivate(ctx context.Context, tx *sql.Tx, userId string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetIsPrivate(ctx, userId)
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
