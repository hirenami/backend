package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateLikeUsecase(ctx context.Context, userId string, tweetId int32) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("tweet does not exist")
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

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.CreateLike(ctx, tx, userId, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.PlusLike(ctx, tx, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	var repid string
	repid, err = u.dao.GetUserId(ctx, tx, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	if err := u.dao.CreateNotification(ctx, tx, userId, repid, "like", tweetId); err != nil {
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

func (u *Usecase) DeleteLikeUsecase(ctx context.Context, userId string, tweetId int32) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("tweet does not exist")
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

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.DeleteLike(ctx, tx, userId, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.MinusLike(ctx, tx, tweetId)
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

func (u *Usecase) CheckLikeUsecase(ctx context.Context, tweetId int32) ([]sqlc.User, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	} else if !bool {
		return nil, errors.New("tweet does not exist")
	}

	likes, err := u.dao.CheckLike(ctx, tx, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	}

	// []Userを作成
	var users []sqlc.User

	// ユーザーIDの配列を展開し、User構造体を取得
	for _, userId := range likes {
		user, err := u.dao.GetProfile(ctx, tx, userId)
		if err != nil {
			// エラーが発生した場合、ロールバック
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}
		users = append(users, user) // ポインタを解 dereferenceしてスライスに追加
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *Usecase) GetUserslikeUsecase(ctx context.Context, userId string) ([]sqlc.Tweet, error) {
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

	tweetid, err := u.dao.GetLikesUser(ctx, tx, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	}

	// []Tweetを作成
	var tweets []sqlc.Tweet

	// ツイートIDの配列を展開し、Tweet構造体を取得
	for _, userId := range tweetid {
		tweet, err := u.dao.GetTweet(ctx, tx, userId)
		if err != nil {
			// エラーが発生した場合、ロールバック
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}
		tweets = append(tweets, tweet) // ポインタを解 dereferenceしてスライスに追加
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tweets, nil
}

func (u *Usecase) IsLikedUsecase(ctx context.Context, userId string, tweetId int32) (bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return false, err
	}

	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return false, err
	} else if !bool {
		return false, errors.New("tweet does not exist")
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

	// daoのメソッドにトランザクションを渡して実行
	bool, err := u.dao.IsLiked(ctx, tx, userId, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return false, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return bool, nil
}