package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) CreateBlock(ctx context.Context, tx *sql.Tx, blockerid string, blockedid string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.CreateBlockParams{
		Blockerid: blockerid,
		Blockedid: blockedid,
	}

	return txQueries.CreateBlock(ctx, arg)
}

func (d *Dao) DeleteBlock(ctx context.Context, tx *sql.Tx, blockerid string, blockedid string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.DeleteBlockParams{
		Blockerid: blockerid,
		Blockedid: blockedid,
	}
	return txQueries.DeleteBlock(ctx, arg)
}

func (d *Dao) IsBlocked(ctx context.Context, tx *sql.Tx, blockerid string, blockedid string) (bool, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	arg := sqlc.IsBlockedParams{
		Blockerid: blockerid,
		Blockedid: blockedid,
	}

	return txQueries.IsBlocked(ctx, arg)
}

func (d *Dao) GetBlocks(ctx context.Context, tx *sql.Tx, blockerid string) ([]string, error) {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	return txQueries.GetBlocks(ctx, blockerid)
}
