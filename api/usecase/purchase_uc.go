package usecase

import (
	"context"
	"fmt"
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