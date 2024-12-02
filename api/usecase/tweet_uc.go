package usecase

import (
	"api/model"
	"context"
	"errors"
	"log"
)

// Tweet メソッドの実装
func (u *Usecase) CreateTweetUsecase(ctx context.Context, myId, content, media_url string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//バリデーション
	if content == ""  && media_url == "" {
		return errors.New("content is empty")
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, myId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.CreateTweet(ctx, tx, myId, content, media_url)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	log.Println(content,media_url)
	//TweetIDを取得
	tweetId, err := u.dao.GetLastInsertID(ctx, tx)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	// ハッシュタグが含まれている場合、ハッシュタグを抽出
	hashtags := ExtractHashtags(content)
	for _, hashtag := range hashtags {
		// ハッシュタグを登録
		err = u.dao.CreateHashtag(ctx, tx, hashtag, int32(tweetId))
		if err != nil {
			// エラーが発生した場合、ロールバック
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return err
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Erase メソッドの実装
func (u *Usecase) EraseTweetUsecase(ctx context.Context, myId string, tweetId int32) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, myId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}

	if _id, err := u.dao.GetUserId(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if _id != myId {
		return errors.New("you can't erase this tweet")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.DeleteTweet(ctx, tx, tweetId)
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

// Edit メソッドの実装
func (u *Usecase) EditTweetUsecase(ctx context.Context, myId string, tweetId int32, content, media_url string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, myId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}

	//バリデーション
	if content == "" && media_url == "" {
		return errors.New("content is empty")
	}
	if media_url != "" && len(media_url) > 255 {
		return errors.New("media_url is too long")
	}

	if _id, err := u.dao.GetUserId(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if _id != myId {
		return errors.New("you can't edit this tweet")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.EditTweet(ctx, tx, content, media_url, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	// ハッシュタグが含まれている場合、ハッシュタグを抽出
	err = u.dao.DeleteHashtag(ctx, tx, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	hashtags := ExtractHashtags(content)
	for _, hashtag := range hashtags {
		// ハッシュタグを登録
		err = u.dao.CreateHashtag(ctx, tx, hashtag, tweetId)
		if err != nil {
			// エラーが発生した場合、ロールバック
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return err
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetUsersTweet メソッドの実装
func (u *Usecase) GetUsersTweetUsecase(ctx context.Context, userId string, myId string) ([]model.TweetParams, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	// ユーザーが存在するかチェック
	exists, err := u.dao.IsUserExists(ctx, tx, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	} else if !exists {
		tx.Rollback()
		return nil, errors.New("user does not exist")
	}

	// ユーザーのツイートを取得
	tweets, err := u.dao.GetUsersTweet(ctx, tx, userId)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	}

	// 結果を格納するためのスライス
	var tweetParamsList []model.TweetParams

	for _, tweet := range tweets {
		tweetParams, err := u.GetTweetParamsUsecase(ctx,tx, myId,int32(tweet.Tweetid))
		if err != nil {
			return nil, err
		}
		if tweetParams != (model.TweetParams{}) {
			tweetParamsList = append(tweetParamsList, tweetParams)
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

// GetTweet メソッドの実装
func (u *Usecase) GetTweetUsecase(ctx context.Context, tweetId int32, myId string) (model.TweetParams, error) {
    // トランザクションを開始
    tx, err := u.dao.Begin()
    if err != nil {
        return model.TweetParams{}, err
    }

    // daoのメソッドにトランザクションを渡して実行
    tweet, err := u.dao.GetTweet(ctx, tx, tweetId)
    if err != nil {
        tx.Rollback()
        return model.TweetParams{}, err
    }

    user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
    if err != nil {
        tx.Rollback()
        return model.TweetParams{}, err
    }

    liked, err := u.dao.IsLiked(ctx, tx, myId, tweet.Tweetid)
    if err != nil {
        tx.Rollback()
        return model.TweetParams{}, err
    }

    retweeted, err := u.dao.IsRetweet(ctx, tx, myId, tweet.Tweetid)
    if err != nil {
        tx.Rollback()
        return model.TweetParams{}, err
    }

	isblocked , err := u.dao.IsBlocked(ctx, tx,tweet.Userid, myId)
	if err != nil {
		tx.Rollback()
		return model.TweetParams{}, err
	}
	log.Println(isblocked,myId,tweet.Userid)

	isfollowing , err := u.dao.IsFollowing(ctx, tx,tweet.Userid,myId)
	if err != nil {
		tx.Rollback()
		return model.TweetParams{}, err
	}
	isprivate := !isfollowing && user.Isprivate && !(myId == user.Userid)

	//impressionをインクリメント
	err = u.dao.PlusImpression(ctx, tx, tweet.Tweetid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return model.TweetParams{}, err
	}

    // トランザクションをコミット
    err = tx.Commit()
    if err != nil {
        return model.TweetParams{}, err
    }

    // 結果をまとめて返す
    return model.TweetParams{
        Tweet:     tweet,
        User:      user,
        Likes:     liked,
        Retweets: retweeted,
		Isblocked: isblocked,
		Isprivate: isprivate,
    }, nil
}