package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Tweet メソッドの実装
func (u *Usecase) CreateTweetUsecase(ctx context.Context, userId, content, media_url string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//バリデーション
	if content == ""  && media_url == "" {
		return errors.New("content is empty")
	}
	if len(content) > 140 {
		return errors.New("content is too long")
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
	err = u.dao.CreateTweet(ctx, tx, userId, content, media_url)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
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
func (u *Usecase) EraseTweetUsecase(ctx context.Context, userId string, tweetId int32) error {
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

	if _id, err := u.dao.GetUserId(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if _id != userId {
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
func (u *Usecase) EditTweetUsecase(ctx context.Context, userId string, tweetId int32, content, media_url string) error {
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

	//バリデーション
	if content == "" {
		return errors.New("content is empty")
	}
	if len(content) > 140 {
		return errors.New("content is too long")
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
	} else if _id != userId {
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
func (u *Usecase) GetUsersTweetUsecase(ctx context.Context, userId string,id string) ([]sqlc.Tweet,[]sqlc.User, []bool ,[]bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil,nil,nil,nil,err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, nil,nil,nil,err
	} else if !bool {
		return nil,nil,nil,nil, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweets, err := u.dao.GetUsersTweet(ctx, tx, userId)
	if err != nil {
		return nil, nil,nil,nil,err
	}

	users := make([]sqlc.User, len(tweets))
	liked := make([]bool, len(tweets))
	retweeted := make([]bool, len(tweets))

	for i, tweet := range tweets {
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			return nil, nil,nil,nil,err
		}
		users[i] = user
	}

	for i, tweet := range tweets {
		bool, err := u.dao.IsLiked(ctx, tx, id, tweet.Tweetid)
		if err != nil {
			return nil, nil,nil,nil,err
		}
		liked[i] = bool
	}

	for i, tweet := range tweets {
		bool, err := u.dao.IsRetweet(ctx, tx, id, tweet.Tweetid)
		if err != nil {
			return nil, nil,nil,nil,err
		}
		retweeted[i] = bool
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, nil,nil,nil,err
	}

	return tweets, users, liked, retweeted, nil
}

// GetTweet メソッドの実装
func (u *Usecase) GetTweetUsecase(ctx context.Context, tweetId int32, id string) (sqlc.Tweet,sqlc.User,bool,bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	}

	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	} else if !bool {
		return sqlc.Tweet{}, sqlc.User{}, false, false, errors.New("tweet does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweet, err := u.dao.GetTweet(ctx, tx, tweetId)
	if err != nil {
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	}

	user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
	if err != nil {
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	}

	liked, err := u.dao.IsLiked(ctx, tx, id, tweet.Tweetid)
	if err != nil {
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	}

	retweeted, err := u.dao.IsRetweet(ctx, tx, id, tweet.Tweetid)
	if err != nil {
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return sqlc.Tweet{}, sqlc.User{}, false, false, err
	}

	return tweet, user, liked, retweeted, nil
}
