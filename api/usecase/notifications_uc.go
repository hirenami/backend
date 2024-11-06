package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) GetNotificationsUsecase(ctx context.Context, userId string) ([]sqlc.Notification,[]sqlc.User,[]sqlc.Tweet, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, nil, nil, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, nil, nil, err
	} else if !bool {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, nil, nil, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	notifications, err := u.dao.GetNotifications(ctx, tx, userId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, nil, nil, err
	}

	// 通知のユーザー情報を取得
	user := make([]sqlc.User, len(notifications))
	tweet := make([]sqlc.Tweet, len(notifications))

	for i, notification := range notifications {
		_user, err := u.dao.GetProfile(ctx, tx, notification.Senderid)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return nil, nil, nil, err
		}
		user[i] = _user
		if(notification.Contentid.Valid){
			_tweet, err := u.dao.GetTweet(ctx, tx, notification.Contentid.Int32)
			if err != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
				}
				return nil, nil, nil, err
			}
			tweet[i] = _tweet
			
		}else{
		 	tweet[i] = sqlc.Tweet{}
		}
		
	}



	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, nil, nil, err
	}

	return notifications, user, tweet, nil
}

// Usecase メソッドの実装
func (u *Usecase) UpdateNotificationStatusUsecase(ctx context.Context, notificationsId int32) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsNotificationExists(ctx, tx, notificationsId); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("notification does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.UpdateNotificationStatus(ctx, tx, notificationsId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
