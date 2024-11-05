package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateReplyUsecase(ctx context.Context, userId, content, media_url string, tweetId int32) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//バリデーション
	if content == "" {
		return errors.New("content is required")
	}
	if len(content) > 140 {
		return errors.New("content is too long")
	}
	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}
	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		return err
	} else if !bool {
		return errors.New("tweet does not exist")
	}

	//Daoのメソッドを呼び出し
	err = u.dao.CreateReply(ctx, tx, userId, content, media_url)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	replyId, err := u.dao.GetLastInsertID(ctx, tx)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	err = u.dao.RelateReplyToTweet(ctx, tx, tweetId, int32(replyId))
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	err = u.dao.PlusReplyCount(ctx, tx, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	repid, err := u.dao.GetUserId(ctx, tx, tweetId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	err = u.dao.CreateNotification(ctx, tx, userId, repid, "reply", int32(replyId))
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

func (u *Usecase) DeleteReplyUsecase(ctx context.Context, userId string, replyId int32) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//バリデーション
	if bool, err := u.dao.IsTweetExists(ctx, tx, replyId); err != nil {
		return err
	} else if !bool {
		return errors.New("reply does not exist")
	}
	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}
	if _id, err := u.dao.GetUserId(ctx, tx, replyId); err != nil {
		return err
	} else if userId != _id {
		return errors.New("you are not the owner of this reply")
	}

	tweetId, err := u.dao.GetTweetRepliedTo(ctx, tx, replyId)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}
	err = u.dao.DeleteTweet(ctx, tx, replyId)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.DeleteRelateReplyToTweet(ctx, tx, tweetId, replyId)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.MinusReplyCount(ctx, tx, tweetId)
	if err != nil {
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

func (u *Usecase) GetUsersReplyUsecase(ctx context.Context, userId string, Id string) ([]sqlc.Tweet, []sqlc.User,[]bool,[]bool,error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, nil,nil,nil,err
	}

	//バリデーション
	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return nil, nil,nil,nil,err
	} else if !bool {
		return nil, nil,nil,nil,errors.New("user does not exist")
	}

	//Daoのメソッドを呼び出し
	tweets, err := u.dao.GetUsersReplies(ctx, tx, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
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
		bool, err := u.dao.IsLiked(ctx, tx, Id, tweet.Tweetid)
		if err != nil {
			return nil, nil,nil,nil,err
		}
		liked[i] = bool
	}

	for i, tweet := range tweets {
		bool, err := u.dao.IsRetweet(ctx, tx, Id, tweet.Tweetid)
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

	return tweets, users,liked,retweeted,nil
}

func (u *Usecase) GetReplyUsecase(ctx context.Context, tweetId int32, Id string) ([]sqlc.Tweet, []sqlc.User, []bool, []bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	//バリデーション
	if bool, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		return nil, nil, nil, nil, err
	} else if !bool {
		return nil, nil, nil, nil, errors.New("reply does not exist")
	}

	//Daoのメソッドを呼び出し
	tweets, err := u.dao.GetRepliesToTweet(ctx, tx, tweetId)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	replies:= make([]sqlc.Tweet, len(tweets))
	users := make([]sqlc.User, len(tweets))
	liked := make([]bool, len(tweets))
	retweeted := make([]bool, len(tweets))

	
	for i, tweet := range tweets {
		reply, err := u.dao.GetTweet(ctx, tx, tweet)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		replies[i] = reply

		user, err := u.dao.GetProfile(ctx, tx, reply.Userid)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		users[i] = user

		bool, err := u.dao.IsLiked(ctx, tx, Id, tweet)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		liked[i] = bool

		bool, err = u.dao.IsRetweet(ctx, tx, Id, tweet)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		retweeted[i] = bool

	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return replies, users, liked, retweeted, nil
}


