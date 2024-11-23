package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) GetPurchase (ctx context.Context,tx *sql.Tx, purchaseId int32) (sqlc.Purchase, error) {
	
	txqueries := d.WithTx(tx)

	return txqueries.GetPurchase(ctx,purchaseId)
}

func (d *Dao) GetPurchasesByUser (ctx context.Context,tx *sql.Tx, userId string) ([]sqlc.Purchase, error) {
	
	txqueries := d.WithTx(tx)

	return txqueries.GetUserPurchases(ctx,userId)
}

func (d *Dao) CreatePurchase (ctx context.Context,tx *sql.Tx, userId string, listingId int32) error {
	
	txqueries := d.WithTx(tx)

	args := sqlc.CreatePurchaseParams{
		Userid: userId,
		Listingid: listingId,
	}

	return txqueries.CreatePurchase(ctx,args)
}