package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateFollowUsecase(ctx context.Context, userId string, followId string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, followId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("follow user does not exist")
	}
	if userId == followId {
		return errors.New("can't follow yourself")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.CreateFollow(ctx, tx, followId, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.CreateNotification(ctx, tx, userId, followId, "follow", 0)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) DelateFollowUsecase(ctx context.Context, userId string, followId string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, followId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("follow user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.DeleteFollow(ctx, tx, followId, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetFollowingUsecase(ctx context.Context, userId string) ([]sqlc.User, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return []sqlc.User{}, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return []sqlc.User{}, err
	} else if !bool {
		return []sqlc.User{},errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	follows, err := u.dao.GetFollowing(ctx, tx, userId)
	if err != nil {
		return []sqlc.User{}, err
	}

	following := []sqlc.User{}

	for _, followId := range follows {
		follow, err := u.dao.GetProfile(ctx, tx, followId)
		if err != nil {
			return []sqlc.User{}, err
		}
		following = append(following, follow)
	}
	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return []sqlc.User{}, err
	}

	return following, nil
}

func (u *Usecase) GetFollowerUsecase(ctx context.Context, userId string) ([]sqlc.User, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return []sqlc.User{}, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return []sqlc.User{}, err
	} else if !bool {
		return []sqlc.User{}, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	followersId, err := u.dao.GetFollower(ctx, tx, userId)
	if err != nil {
		return []sqlc.User{}, err
	}

	followers := []sqlc.User{}

	for _, followerId := range followersId {
		follower, err := u.dao.GetProfile(ctx, tx, followerId)
		if err != nil {
			return []sqlc.User{}, err
		}
		followers = append(followers, follower)
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return []sqlc.User{}, err
	}

	return followers, nil
}

func (u *Usecase) GetFollowCountUsecase(ctx context.Context, userId string) (int32, int32, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return 0, 0, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return 0, 0, err
	} else if !bool {
		return 0, 0, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	following, err := u.dao.CountFollowing(ctx, tx, userId)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return 0, 0, err
		}
	}

	followers, err := u.dao.CountFollower(ctx, tx, userId)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return 0, 0, err
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return 0, 0, err
	}

	return int32(following), int32(followers), nil
}

func (u *Usecase) IsFollowingUsecase(ctx context.Context, userId string, followId string) (bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return false, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return false, err
	} else if !bool {
		return false, errors.New("user does not exist")
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, followId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return false, err
	} else if !bool {
		return false, errors.New("follow user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	isFollowing, err := u.dao.IsFollowing(ctx, tx, followId, userId)
	if err != nil {
		return false, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return isFollowing, nil
}