package usecase

import (
	"api/sqlc"
	"api/model"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) SearchByKeywordUsecase(ctx context.Context, keyword string) ([]model.TweetParams, error) {
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

	// キーワードが空の場合はエラー
	if keyword == "" {
		return nil, errors.New("keyword is required")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweets, err := u.dao.SearchByKeyword(ctx, tx, keyword)
	if err != nil {
		return nil, err
	}

	// TweetParamsのスライスを作成
	var tweetParamsList []model.TweetParams

	for _, tweet := range tweets {
		// ユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			return nil, err
		}

		// いいねとリツイート情報を取得
		liked, err := u.dao.IsLiked(ctx, tx, tweet.Userid, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		retweeted, err := u.dao.IsRetweet(ctx, tx, tweet.Userid, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		// TweetParamsに追加
		tweetParamsList = append(tweetParamsList, model.TweetParams{
			Tweet:   tweet,
			User:    user,
			Likes:   liked,
			Retweets: retweeted,
		})
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}


// Usecase メソッドの実装
func (u *Usecase) SearchByUserUsecase(ctx context.Context, keyword string) ([]sqlc.User,error) {
	
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil,err
	}

	if keyword == "" {
		return nil,errors.New("keyword is required")
	}

	// daoのメソッドにトランザクションを渡して実行
	users,err := u.dao.SearchUser(ctx, tx, keyword)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil,err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil,err
	}

	return users,nil
}

// Usecase メソッドの実装
func (u *Usecase) SearchByHashtagUsecase(ctx context.Context, keyword string) ([]model.TweetParams, error) {
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

	// キーワードが空の場合はエラー
	if keyword == "" {
		return nil, errors.New("keyword is required")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweets, err := u.dao.SearchByHashtag(ctx, tx, keyword)
	if err != nil {
		return nil, err
	}

	// TweetParamsのスライスを作成
	var tweetParamsList []model.TweetParams

	for _, tweet := range tweets {
		// ユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			return nil, err
		}

		// いいねとリツイート情報を取得
		liked, err := u.dao.IsLiked(ctx, tx, tweet.Userid, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		retweeted, err := u.dao.IsRetweet(ctx, tx, tweet.Userid, tweet.Tweetid)
		if err != nil {
			return nil, err
		}

		// TweetParamsに追加
		tweetParamsList = append(tweetParamsList, model.TweetParams{
			Tweet:   tweet,
			User:    user,
			Likes:   liked,
			Retweets: retweeted,
		})
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}