package usecase

import (
	"api/model"
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

func (u *Usecase) GetUserslikeUsecase(ctx context.Context, myId, userId string) ([]model.TweetParams, error) {
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

	// 結果を格納するためのスライス
	tweetParamsList := make([]model.TweetParams, len(tweetid))

	for i, tweet := range tweetid {
		// ツイート情報を取得
		tweet, err := u.dao.GetTweet(ctx, tx, tweet)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}

		// ユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}

		// ツイートが「いいね」されているか確認
		liked, err := u.dao.IsLiked(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}

		// ツイートがリツイートされているか確認
		retweeted, err := u.dao.IsRetweet(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}

		isblocked, err := u.dao.IsBlocked(ctx, tx, myId, tweet.Userid)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		isfollowing, err := u.dao.IsFollowing(ctx, tx, tweet.Userid, myId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		isprivate := !isfollowing && user.Isprivate

		// TweetParams構造体にデータをまとめる
		tweetParamsList[i] = model.TweetParams{
			Tweet:     tweet,
			User:      user,
			Likes:     liked,
			Retweets:  retweeted,
			Isblocked: isblocked,
			Isprivate: isprivate,
		}

		//impressionをインクリメント
		err = u.dao.PlusImpression(ctx, tx, tweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, err
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}
