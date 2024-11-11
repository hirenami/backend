package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateBiography(ctx context.Context, tx *sql.Tx, userId string, biography string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateBiographyParams{
		Userid: userId,
		Biography: biography,
	}
	return txQueries.CreateBiography(ctx, arg)
}

func (d* Dao) CreateHeaderImage(ctx context.Context, tx *sql.Tx, userId string, header_image string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateHeaderImageParams{
		Userid: userId,
		HeaderImage: header_image,
	}
	return txQueries.CreateHeaderImage(ctx, arg)
}

func (d *Dao) CreateIconImage(ctx context.Context, tx *sql.Tx, userId string, icon_image string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateIconImageParams{
		Userid: userId,
		IconImage: icon_image,
	}
	return txQueries.CreateIconImage(ctx, arg)
}

func (d *Dao) GetProfile(ctx context.Context, tx *sql.Tx, userId string) (sqlc.User, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetProfile(ctx, userId)
}