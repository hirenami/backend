package usecase

import (
	"api/model"
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

func (u *Usecase) GetFollowingUsecase(ctx context.Context, myId, userId string) ([]model.Profile, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	} else if !bool {
		return nil, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	follows, err := u.dao.GetFollowing(ctx, tx, userId)
	if err != nil {
		return nil, err
	}

	// 結果を格納するためのスライス
	followsParamsList := make([]model.Profile, len(follows))

	for i, followId := range follows {
		user, err := u.dao.GetProfile(ctx, tx, followId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isFollowing, err := u.dao.IsFollowing(ctx, tx, followId, myId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isFollower, err := u.dao.IsFollowing(ctx, tx, myId, followId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		following, err := u.dao.CountFollowing(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		followers, err := u.dao.CountFollower(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		// Params構造体にデータをまとめる
		followsParamsList[i] = model.Profile{
			User:        user,
			Follows:     int32(following),
			Followers:   int32(followers),
			Isfollows:   isFollowing,
			Isfollowers: isFollower,
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return followsParamsList, nil
}

func (u *Usecase) GetFollowerUsecase(ctx context.Context, userId, myId string) ([]model.Profile, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	} else if !bool {
		return nil, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	followers, err := u.dao.GetFollower(ctx, tx, userId)
	if err != nil {
		return nil, err
	}

	// 結果を格納するためのスライス
	followsParamsList := make([]model.Profile, len(followers))

	for i, followersId := range followers {
		user, err := u.dao.GetProfile(ctx, tx, followersId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isFollowing, err := u.dao.IsFollowing(ctx, tx, followersId, myId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isFollower, err := u.dao.IsFollowing(ctx, tx, myId, followersId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		following, err := u.dao.CountFollowing(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		followers, err := u.dao.CountFollower(ctx, tx, userId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		// Params構造体にデータをまとめる
		followsParamsList[i] = model.Profile{
			User:        user,
			Follows:     int32(following),
			Followers:   int32(followers),
			Isfollows:   isFollowing,
			Isfollowers: isFollower,
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return followsParamsList, nil
}
