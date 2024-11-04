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

func (u *Usecase) GetUsersReply(ctx context.Context, userId string) ([]sqlc.Tweet, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	//バリデーション
	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return nil, err
	} else if !bool {
		return nil, errors.New("user does not exist")
	}

	//Daoのメソッドを呼び出し
	tweets, err := u.dao.GetUsersReplies(ctx, tx, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
