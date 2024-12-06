package usecase

import (
	"api/dao"
	"api/model"
	"context"
	"database/sql"
	"log"
)

// Usecase構造体
type Usecase struct {
	dao *dao.Dao
}

// NewTestUsecase コンストラクタ
func NewUsecase(dao *dao.Dao) *Usecase {
	return &Usecase{
		dao: dao,
	}
}

func (u *Usecase) GetTweetParamsUsecase(ctx context.Context, tx *sql.Tx, myId string, tweetid int32) (model.TweetParams, error) {

		// ツイート情報を取得
		tweet, err := u.dao.GetTweet(ctx, tx, tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParams{}, err
		}

		// ユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParams{}, err
		}

		var retweet model.TweetParam
		if(tweet.Retweetid != 0) {
		retweet, err = u.GetTweetParamUsecase(ctx, tx, myId, tweet.Retweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParams{}, err
		}
		}else{
			retweet = model.TweetParam{}
		}

		// ツイートが「いいね」されているか確認
		liked, err := u.dao.IsLiked(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParams{}, err
		}

		// ツイートがリツイートされているか確認
		retweeted, err := u.dao.IsRetweet(ctx, tx, myId, tweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParams{}, err
		}

		isblocked, err := u.dao.IsBlocked(ctx, tx, tweet.Userid, myId)
		if err != nil {
			tx.Rollback()
			return model.TweetParams{}, err
		}

		isfollowing, err := u.dao.IsFollowing(ctx, tx, tweet.Userid, myId)
		if err != nil {
			tx.Rollback()
			return model.TweetParams{}, err
		}
		isprivate := !isfollowing && user.Isprivate && !(myId == tweet.Userid)

		if isblocked || isprivate || tweet.Isdeleted {
			return model.TweetParams{}, nil
		}

		// TweetParams構造体にデータをまとめる
		tweetParamList := model.TweetParams{
			Tweet:     tweet,
			User:      user,
			Likes:     liked,
			Retweet:  retweet,
			Retweets:  retweeted,
			Isblocked: isblocked,
			Isprivate: isprivate,
		}
	
	return tweetParamList, nil
}

func (u *Usecase) GetTweetParamUsecase(ctx context.Context, tx *sql.Tx, myId string, tweetid int32) (model.TweetParam, error) {

	// ツイート情報を取得
	tweet, err := u.dao.GetTweet(ctx, tx, tweetid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return model.TweetParam{}, err
	}

	// ユーザー情報を取得
	user, err := u.dao.GetProfile(ctx, tx, tweet.Userid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return model.TweetParam{}, err
	}

	var quote model.QuoteParam
	if(tweet.Isquote){
		quotetweet, err := u.dao.GetTweet(ctx, tx, tweet.Retweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParam{}, err
		}
		quoteuser, err := u.dao.GetProfile(ctx, tx, quotetweet.Userid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParam{}, err
		}
		quoteislike ,err := u.dao.IsLiked(ctx, tx, myId, quotetweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParam{}, err
		}
		quoteretweeted, err := u.dao.IsRetweet(ctx, tx, myId, quotetweet.Tweetid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return model.TweetParam{}, err
		}
		isquoteblocked, err := u.dao.IsBlocked(ctx, tx, quotetweet.Userid, myId)
		if err != nil {
			tx.Rollback()
			return model.TweetParam{}, err
		}
		isquotefollowing, err := u.dao.IsFollowing(ctx, tx, quotetweet.Userid, myId)
		if err != nil {
			tx.Rollback()
			return model.TweetParam{}, err
		}
		isquoteprivate := !isquotefollowing && quoteuser.Isprivate && !(myId == quotetweet.Userid)
		if isquoteblocked || isquoteprivate || quotetweet.Isdeleted {
			return model.TweetParam{}, nil
		}
		quote = model.QuoteParam{
			Tweet:     quotetweet,
			User:      quoteuser,
			Likes:     quoteislike,
			Retweets:  quoteretweeted,
			Isblocked: isquoteblocked,
			Isprivate: isquoteprivate,
		}
	}else{
		quote = model.QuoteParam{}
	}

	// ツイートが「いいね」されているか確認
	liked, err := u.dao.IsLiked(ctx, tx, myId, tweet.Tweetid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return model.TweetParam{}, err
	}

	// ツイートがリツイートされているか確認
	retweeted, err := u.dao.IsRetweet(ctx, tx, myId, tweet.Tweetid)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return model.TweetParam{}, err
	}

	isblocked, err := u.dao.IsBlocked(ctx, tx, tweet.Userid, myId)
	if err != nil {
		tx.Rollback()
		return model.TweetParam{}, err
	}

	isfollowing, err := u.dao.IsFollowing(ctx, tx, tweet.Userid, myId)
	if err != nil {
		tx.Rollback()
		return model.TweetParam{}, err
	}
	isprivate := !isfollowing && user.Isprivate && !(myId == tweet.Userid)

	// TweetParams構造体にデータをまとめる
	tweetParamList := model.TweetParam{
		Tweet:     tweet,
		User:      user,
		Quote:    quote,
		Likes:     liked,
		Retweets:  retweeted,
		Isblocked: isblocked,
		Isprivate: isprivate,
	}

return tweetParamList, nil
}