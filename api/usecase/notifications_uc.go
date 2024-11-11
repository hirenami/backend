package usecase

import (
	"api/sqlc"
	"api/model"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) GetNotificationsUsecase(ctx context.Context, myId string) ([]model.NotificationParams, error) {
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
	exists, err := u.dao.IsUserExists(ctx, tx, myId)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("user does not exist")
	}

	// 通知の取得
	notifications, err := u.dao.GetNotifications(ctx, tx, myId)
	if err != nil {
		return nil, err
	}

	// 通知のユーザー情報とツイート情報を取得
	notificationParams := make([]model.NotificationParams, len(notifications))
	for i, notification := range notifications {
		// ユーザー情報の取得
		_user, err := u.dao.GetProfile(ctx, tx, notification.Senderid)
		if err != nil {
			return nil, err
		}

		// ツイート情報の取得（ContentidがValidな場合のみ）
		var _tweet sqlc.Tweet
		if notification.Contentid!=0 {
			_tweet, err = u.dao.GetTweet(ctx, tx, notification.Contentid)
			if err != nil {
				return nil, err
			}
		}

		// NotificationParamsに詰め込む
		notificationParams[i] = model.NotificationParams{
			Notification: notification,
			User:         _user,
			Tweet:        _tweet,
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// 最終結果を返す
	return notificationParams, nil
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
