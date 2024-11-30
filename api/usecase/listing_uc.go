package usecase

import (
	"context"
	"api/model"
)

func (u* Usecase) CreateListingUsecase(ctx context.Context,listingId int64, userId string,content string , media_url string, listing model.Listing) error {
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	err = u.dao.CreateTweet(ctx, tx,userId, content, media_url)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
	}

	tweetId , err := u.dao.GetLastInsertID(ctx, tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
	}

	err = u.dao.CreateListing(ctx, tx, listingId, userId, int32(tweetId), listing.Listingname, listing.Listingdescription, listing.Listingprice, listing.Type, listing.Stock, listing.Condition)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	}

	err = u.dao.UpdateReview(ctx, tx, int32(tweetId),-1)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (u* Usecase) GetListingUsecase(ctx context.Context,myId string, listingid int64) (model.ListingDetails, error) {
	tx , err := u.dao.Begin()
	if err != nil {
		return model.ListingDetails{}, err
	}

	listing , err := u.dao.GetListing(ctx, tx, listingid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return model.ListingDetails{}, err
		}
		return model.ListingDetails{}, err
	}

	tweet , err := u.dao.GetTweet(ctx, tx, int32(listing.Tweetid))
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return model.ListingDetails{}, err
		}
		return model.ListingDetails{}, err
	}

	userIds , err := u.dao.GetPurchaseByListingId(ctx,tx, listing.Listingid)
	if err != nil {
		return model.ListingDetails{}, err
	}

	users := make([]model.Profile, len(userIds))

	for i, userId := range userIds {
		user, err := u.dao.GetProfile(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return model.ListingDetails{}, err
			}
		}

		isFollowing, err := u.dao.IsFollowing(ctx, tx, userId, myId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return model.ListingDetails{}, err
			}
		}

		isFollower, err := u.dao.IsFollowing(ctx, tx, myId, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return model.ListingDetails{}, err
			}
		}

		following, err := u.dao.CountFollowing(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return model.ListingDetails{}, err
			}
		}

		followers, err := u.dao.CountFollower(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return model.ListingDetails{}, err
			}
		}

		isblocked , err := u.dao.IsBlocked(ctx, tx, myId, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return model.ListingDetails{}, err
			}
		}
		if isblocked {
			continue
		}

		isprivate := !isFollowing && user.Isprivate

		// Params構造体にデータをまとめる
		users[i] = model.Profile{
			User:        user,
			Follows:     int32(following),
			Followers:   int32(followers),
			Isfollows:   isFollowing,
			Isfollowers: isFollower,
			Isblocked:   isblocked,
			Isprivate:  isprivate,
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return model.ListingDetails{}, err
	}

	return model.ListingDetails{
		Listing: listing,
		User: users,
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