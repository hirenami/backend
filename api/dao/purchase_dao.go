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

func (d *Dao) CreatePurchase (ctx context.Context,tx *sql.Tx, userId string, listingId int64) error {
	
	txqueries := d.WithTx(tx)

	args := sqlc.CreatePurchaseParams{
		Userid: userId,
		Listingid: listingId,
	}

	return txqueries.CreatePurchase(ctx,args)
}

func (d *Dao) GetPurchaseByListingId (ctx context.Context,tx *sql.Tx, purchaseId int64) ([]string,error) {
	
	txqueries := d.WithTx(tx)

	return txqueries.GetPurchaseByListing(ctx,purchaseId)
}