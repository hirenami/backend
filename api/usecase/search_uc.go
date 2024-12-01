package usecase

import (
	"api/model"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) SearchByKeywordUsecase(ctx context.Context,myId, keyword string) ([]model.TweetParams, error) {
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
		tweetParams, err := u.GetTweetParamsUsecase(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			return nil, err
		}
		if(tweetParams!=model.TweetParams{}){
			tweetParamsList = append(tweetParamsList, tweetParams)
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}

// Usecase メソッドの実装
func (u *Usecase) SearchByUserUsecase(ctx context.Context, myId, keyword string) ([]model.Profile, error) {

	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	if keyword == "" {
		return nil, errors.New("keyword is required")
	}

	// daoのメソッドにトランザクションを渡して実行
	users, err := u.dao.SearchUser(ctx, tx, keyword)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	}

	// 結果を格納するためのスライス
	UserParamsList := make([]model.Profile, len(users))

	for i, user := range users {
		// ユーザー情報を取得
		isFollowing, err := u.dao.IsFollowing(ctx, tx, user.Userid, myId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isFollower, err := u.dao.IsFollowing(ctx, tx, myId, user.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		following, err := u.dao.CountFollowing(ctx, tx, user.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		followers, err := u.dao.CountFollower(ctx, tx, user.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isblocked , err := u.dao.IsBlocked(ctx, tx, user.Userid, myId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		if(isblocked){
			continue
		}

		isprivate := !isFollowing && user.Isprivate && !(myId == user.Userid)

		isblock, err := u.dao.IsBlocked(ctx, tx, myId,user.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		isrequest, err := u.dao.IsKeyFollowExists(ctx, tx, user.Userid,myId)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, err
			}
		}

		// Params構造体にデータをまとめる
		UserParamsList[i] = model.Profile{
			User:        user,
			Follows:     int32(following),
			Followers:   int32(followers),
			Isfollows:   isFollowing,
			Isfollowers: isFollower,
			Isblocked:   isblocked,
			Isprivate:   isprivate,
			Isblock:     isblock,
			Isrequest:   isrequest,
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return UserParamsList, nil
}

// Usecase メソッドの実装
func (u *Usecase) SearchByHashtagUsecase(ctx context.Context, myId,keyword string) ([]model.TweetParams, error) {
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
		tweetParams, err := u.GetTweetParamsUsecase(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			return nil, err
		}
		if(tweetParams!=model.TweetParams{}){
			tweetParamsList = append(tweetParamsList, tweetParams)
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tweetParamsList, nil
}
