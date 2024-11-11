package usecase

import (
	"api/model"
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

func (u *Usecase) GetUsersReplyUsecase(ctx context.Context, userId string, myId string) ([]model.TweetParams, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		// エラーが発生した場合はロールバック
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
		}
	}()

	// ユーザーが存在するかをバリデーション
	if exists, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("user does not exist")
	}

	// ユーザーのリプライツイートを取得
	tweets, err := u.dao.GetUsersReplies(ctx, tx, userId)
	if err != nil {
		return nil, err
	}

	// TweetParamsのスライスを作成
	var tweetParamsList []model.TweetParams

	for _, tweet := range tweets {
		// ツイート投稿者のユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			return nil, err
		}

		// いいねとリツイート情報を取得
		liked, err := u.dao.IsLiked(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		retweeted, err := u.dao.IsRetweet(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		// TweetParamsに情報を詰めてリストに追加
		tweetParamsList = append(tweetParamsList, model.TweetParams{
			Tweet:    tweet,
			User:     user,
			Likes:    liked,
			Retweets: retweeted,
		})
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}

func (u *Usecase) GetReplyUsecase(ctx context.Context, tweetId int32, myId string) ([]model.TweetParams, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		// エラーが発生した場合はロールバック
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
		}
	}()

	// ツイートが存在するかバリデーション
	if exists, err := u.dao.IsTweetExists(ctx, tx, tweetId); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("reply does not exist")
	}

	// ツイートへのリプライを取得
	tweets, err := u.dao.GetRepliesToTweet(ctx, tx, tweetId)
	if err != nil {
		return nil, err
	}

	// TweetParamsのスライスを作成
	var tweetParamsList []model.TweetParams

	for _, tweet := range tweets {
		// リプライツイートを取得
		reply, err := u.dao.GetTweet(ctx, tx, tweet)
		if err != nil {
			return nil, err
		}

		// ツイート投稿者のユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, reply.Userid)
		if err != nil {
			return nil, err
		}

		// いいねとリツイート情報を取得
		liked, err := u.dao.IsLiked(ctx, tx, myId, tweet)
		if err != nil {
			return nil, err
		}

		retweeted, err := u.dao.IsRetweet(ctx, tx, myId, tweet)
		if err != nil {
			return nil, err
		}

		// TweetParamsに情報を詰めてリストに追加
		tweetParamsList = append(tweetParamsList, model.TweetParams{
			Tweet:    reply,
			User:     user,
			Likes:    liked,
			Retweets: retweeted,
		})
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}

func (u *Usecase) GetTweetRepliedToUsecase(ctx context.Context, Id string, replyId int32) ([]model.TweetParams, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		// エラーが発生した場合はロールバック
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
		}
	}()

	// バリデーション: リプライの存在確認
	if exists, err := u.dao.IsTweetExists(ctx, tx, replyId); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("reply does not exist")
	}

	// ツイートのリストを初期化
	var tweets []sqlc.Tweet

	// リプライツイートを辿る
	for {
		// リプライが存在するか確認
		exists, err := u.dao.IsReplyExists(ctx, tx, replyId)
		if err != nil {
			return nil, err
		}
		if !exists {
			break
		}

		// リプライ元ツイートのIDを取得
		tweetId, err := u.dao.GetTweetRepliedTo(ctx, tx, replyId)
		if err != nil {
			return nil, err
		}
		
		// ツイートを取得
		_tweet, err := u.dao.GetTweet(ctx, tx, tweetId)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, _tweet)

		// 次のリプライ元IDに更新
		replyId = tweetId
	}

	// TweetParamsにまとめるためのリストを準備
	var tweetParamsList []model.TweetParams
	users := make([]sqlc.User, len(tweets))
	liked := make([]bool, len(tweets))
	retweeted := make([]bool, len(tweets))

	// ツイートごとにユーザー、いいね、リツイート情報を取得
	for i, tweet := range tweets {
		// ツイートの投稿者情報を取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			return nil, err
		}
		users[i] = user

		// いいね情報を取得
		isLiked, err := u.dao.IsLiked(ctx, tx, Id, tweet.Tweetid)
		if err != nil {
			return nil, err
		}
		liked[i] = isLiked

		// リツイート情報を取得
		isRetweeted, err := u.dao.IsRetweet(ctx, tx, Id, tweet.Tweetid)
		if err != nil {
			return nil, err
		}
		retweeted[i] = isRetweeted

		// TweetParamsとしてまとめる
		tweetParamsList = append(tweetParamsList, model.TweetParams{
			Tweet:    tweet,
			User:     user,
			Likes:    isLiked,
			Retweets: isRetweeted,
		})
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}



