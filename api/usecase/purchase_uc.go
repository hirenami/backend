package usecase

import (
	"context"
	"fmt"
	"api/model"
)

func (u * Usecase) UpdatePremiumUsecase (ctx context.Context, myId string) error {
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, myId); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	} else if !bool {
		return fmt.Errorf("user does not exist")
	}

	err = u.dao.CreateIsPremium(ctx, tx, myId)
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

func (u *Usecase) GetPurchaseUsecase(ctx context.Context, purchaseId int32) (model.PurchaseParams, error) {
	tx , err := u.dao.Begin()
	if err != nil {
		return model.PurchaseParams{}, err
	}
	purchase , err := u.dao.GetPurchase(ctx, tx, purchaseId)
	if err != nil {
		return model.PurchaseParams{}, err
	}
	listing , err := u.dao.GetListing(ctx, tx, purchase.Listingid)
	if err != nil {
		return model.PurchaseParams{}, err
	}
	user , err := u.dao.GetProfile(ctx, tx, listing.Userid)
	if err != nil {
		return model.PurchaseParams{}, err
	}
	tweet , err := u.dao.GetTweet(ctx, tx, listing.Tweetid)
	if err != nil {
		return model.PurchaseParams{}, err
	}
	if err := tx.Commit(); err != nil {
		return model.PurchaseParams{}, err
	}
	return model.PurchaseParams{
		Purchase: purchase,
		Listing: listing,
		User: user,
		Tweet: tweet,
	}, nil
}

func (u *Usecase) GetPurchasesByUserUsecase(ctx context.Context, userId string) ([]model.PurchaseParams, error) {
	tx , err := u.dao.Begin()
	if err != nil {
		return []model.PurchaseParams{}, err
	}
	purchases , err := u.dao.GetPurchasesByUser(ctx, tx, userId)
	if err != nil {
		return []model.PurchaseParams{}, err
	}
	var purchaseParams []model.PurchaseParams
	for _, purchase := range purchases {
		listing , err := u.dao.GetListing(ctx, tx, purchase.Listingid)
		if err != nil {
			return []model.PurchaseParams{}, err
		}
		user , err := u.dao.GetProfile(ctx, tx, listing.Userid)
		if err != nil {
			return []model.PurchaseParams{}, err
		}
		tweet , err := u.dao.GetTweet(ctx, tx, listing.Tweetid)
		if err != nil {
			return []model.PurchaseParams{}, err
		}
		purchaseParams = append(purchaseParams, model.PurchaseParams{
			Purchase: purchase,
			Listing: listing,
			User: user,
			Tweet: tweet,
		})
	}
	if err := tx.Commit(); err != nil {
		return []model.PurchaseParams{}, err
	}
	return purchaseParams, nil
}

func (u *Usecase) CreatePurchaseUsecase(ctx context.Context, userId string, listingId int64) error {
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	} else if !bool {
		return fmt.Errorf("user does not exist")
	}

	err = u.dao.CreatePurchase(ctx, tx, userId, listingId)
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