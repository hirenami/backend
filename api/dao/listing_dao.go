package dao

import (
	"api/sqlc"
	"database/sql"
	"context"
)

func (d *Dao) GetListing(ctx context.Context,tx *sql.Tx, listingid int64) (sqlc.Listing, error) {
	
	txQueries := d.WithTx(tx)

	return txQueries.GetListing(ctx, listingid)
}

func (d *Dao) GetListingByTweet(ctx context.Context,tx *sql.Tx, tweetid int32) (sqlc.Listing, error) {
	
	txQueries := d.WithTx(tx)

	return txQueries.GetListingByTweet(ctx, tweetid)
}

func (d *Dao) CreateListing(ctx context.Context,tx *sql.Tx,listingid int64, userid string, tweetid int32, listingname string, listingdescription string, listingprice int32, Type string,stock int32,condition string) error {
	
	txQueries := d.WithTx(tx)

	return txQueries.CreateListing(ctx, sqlc.CreateListingParams{
		Listingid: listingid,
		Userid: userid,
		Tweetid: tweetid,
		Listingname: listingname,
		Listingdescription: listingdescription,
		Listingprice: listingprice,
		Type: Type,
		Stock: stock,
		Condition: condition,
	})
}

func (d *Dao) GetUserListings(ctx context.Context,tx *sql.Tx, userid string) ([]sqlc.Listing, error) {
	
	txQueries := d.WithTx(tx)

	return txQueries.GetUserListings(ctx, userid)
}

func (d *Dao) DeleteStock (ctx context.Context,tx *sql.Tx, listingid int64) error {
	
	txQueries := d.WithTx(tx)

	return txQueries.DeleteStock(ctx, listingid)
}

func (d *Dao) UpdateListing (ctx context.Context,tx *sql.Tx, userId string) error {
	
	txQueries := d.WithTx(tx)

	return txQueries.UpdateListing(ctx, userId)
}

func (d *Dao) GetRandomListings (ctx context.Context,tx *sql.Tx) ([]int64, error) {
	
	txQueries := d.WithTx(tx)

	return txQueries.GetRandomListings(ctx)
}

