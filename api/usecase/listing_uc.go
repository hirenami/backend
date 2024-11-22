package usecase

import (
	"context"
	"api/model"
)

func (u* Usecase) CreateListingUsecase(ctx context.Context, listing model.Listing) error {
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	err = u.dao.CreateListing(ctx, tx, listing.Userid, listing.Tweetid, listing.Listingname, listing.Listingdescription, listing.Listingprice, listing.Type, listing.Stock)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (u* Usecase) GetListingUsecase(ctx context.Context, listingid int32) (model.ListingParams, error) {
	tx , err := u.dao.Begin()
	if err != nil {
		return model.ListingParams{}, err
	}
	listing , err := u.dao.GetListing(ctx, tx, listingid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return model.ListingParams{}, err
		}
		return model.ListingParams{}, err
	}
	
	user , err := u.dao.GetProfile(ctx, tx, listing.Userid)
	if err != nil {
		return model.ListingParams{}, err
	}

	tweet, err := u.dao.GetTweet(ctx, tx, listing.Tweetid)
	if err != nil {
		return model.ListingParams{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.ListingParams{}, err
	}

	return model.ListingParams{
		Listing: listing,
		User: user,
		Tweet: tweet,
	}, nil
}

func (u* Usecase) GetListingByTweetUsecase(ctx context.Context, tweetid int32) (model.ListingParams, error) {
	tx , err := u.dao.Begin()
	if err != nil {
		return model.ListingParams{}, err
	}
	listing , err := u.dao.GetListingByTweet(ctx, tx, tweetid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return model.ListingParams{}, err
		}
		return model.ListingParams{}, err
	}
	
	user , err := u.dao.GetProfile(ctx, tx, listing.Userid)
	if err != nil {
		return model.ListingParams{}, err
	}

	tweet, err := u.dao.GetTweet(ctx, tx, listing.Tweetid)
	if err != nil {
		return model.ListingParams{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.ListingParams{}, err
	}

	return model.ListingParams{
		Listing: listing,
		User: user,
		Tweet: tweet,
	}, nil
}

func (u* Usecase) GetUserListingsUsecase(ctx context.Context, userid string) ([]model.ListingParams, error) {
	tx , err := u.dao.Begin()
	if err != nil {
		return []model.ListingParams{}, err
	}
	listings , err := u.dao.GetUserListings(ctx, tx, userid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return []model.ListingParams{}, err
		}
		return []model.ListingParams{}, err
	}
	
	var res []model.ListingParams

	for _, listing := range listings {
		user , err := u.dao.GetProfile(ctx, tx, listing.Userid)
		if err != nil {
			return []model.ListingParams{}, err
		}

		tweet, err := u.dao.GetTweet(ctx, tx, listing.Tweetid)
		if err != nil {
			return []model.ListingParams{}, err
		}

		res = append(res, model.ListingParams{
			Listing: listing,
			User: user,
			Tweet: tweet,
		})
	}

	if err := tx.Commit(); err != nil {
		return []model.ListingParams{}, err
	}

	return res, nil
}