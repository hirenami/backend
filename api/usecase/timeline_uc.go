package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) GetTimelineUsecase(ctx context.Context, Id string) ([]sqlc.Tweet,[]sqlc.User,[]bool,[]bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, nil,nil,nil,err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, Id); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, nil,nil,nil,err
	} else if !bool {
		return nil, nil,nil,nil,errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweets, err := u.dao.Timeline(ctx, tx, Id)
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
