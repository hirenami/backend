package usecase

import (
	"api/model"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) GetTimelineUsecase(ctx context.Context, id string) ([]model.TweetParams, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		// エラーハンドリングのためにトランザクションをロールバックする
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
		}
	}()

	// ユーザー存在確認
	exists, err := u.dao.IsUserExists(ctx, tx, id)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("user does not exist")
	}

	// タイムラインのツイートを取得
	tweets, err := u.dao.Timeline(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	// 結果を格納するスライス
	tweetParams := make([]model.TweetParams, len(tweets))

	// ユーザー情報、いいね、リツイートの情報を取得
	for i, tweet := range tweets {
		// ユーザー情報の取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			return nil, err
		}

		// いいね情報の取得
		liked, err := u.dao.IsLiked(ctx, tx, id, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		// リツイート情報の取得
		retweeted, err := u.dao.IsRetweet(ctx, tx, id, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		// TweetParamsにデータを詰め込む
		tweetParams[i] = model.TweetParams{
			Tweet:     tweet,
			User:      user,
			Likes:     liked,
			Retweets:  retweeted,
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// 最終結果を返す
	return tweetParams, nil
}
