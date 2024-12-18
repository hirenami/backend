package usecase

import (
	"api/model"
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateDm(ctx context.Context, userId, repid, content, media_url string) error {

	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return err
	} else if !bool {
		return errors.New("user not found")
	}
	if bool, err := u.dao.IsUserExists(ctx, tx, repid); err != nil {
		return err
	} else if !bool {
		return errors.New("user not found")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.CreateDm(ctx, tx, userId, repid, content, media_url)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	if content != "" || media_url != "" {
		err = u.dao.CreateNotification(ctx, tx, userId, repid, "dm", 0)
		if err != nil {
			// エラーが発生した場合、ロールバック
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
			}
			return err
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

// Usecase メソッドの実装
func (u *Usecase) GetDms(ctx context.Context, userId, repid string) ([]sqlc.Dm, error) {

	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		return nil, err
	} else if !bool {
		return nil, errors.New("user not found")
	}

	// daoのメソッドにトランザクションを渡して実行
	dms, err := u.dao.GetDms(ctx, tx, userId, repid)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	}

	err = u.dao.SetDmStatus(ctx, tx, repid, userId)
	if err != nil {
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

	return dms, nil
}

// Usecase メソッドの実装
func (u *Usecase) DeleteDm(ctx context.Context, dmsid int32) error {

	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.DeleteDm(ctx, tx, dmsid)
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

	return err
}

func (u *Usecase) GetAllDms(ctx context.Context, myId string) ([]model.Conversation, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // コミットされていなければロールバック

	// ユーザーの存在確認
	if exists, err := u.dao.IsUserExists(ctx, tx, myId); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("user not found")
	}

	// 自分が関係するメッセージを取得
	dms, err := u.dao.GetAllDms(ctx, tx, myId)
	if err != nil {
		return nil, err
	}

	// 他ユーザーごとにメッセージを整理
	conversationMap := make(map[string]model.Conversation)
	for _, dm := range dms {
		// 自分以外のユーザー ID を特定
		partnerId := dm.Senderid
		if dm.Senderid == myId {
			partnerId = dm.Receiverid
		}

		// ユーザー情報を取得
		user, err := u.dao.GetProfile(ctx, tx, partnerId)
		if err != nil {
			return nil, err
		}

		bool, err := u.dao.IsBlocked(ctx, tx, partnerId, myId)
		if err != nil {
			return nil, err
		} else if bool {
			continue
		}

		// 既存のデータを取得、なければ新規作成
		conv, exists := conversationMap[partnerId]
		if !exists {
			conv = model.Conversation{
				User: user,
				Dms:  []sqlc.Dm{},
			}
		}

		// DM を追加
		conv.Dms = append(conv.Dms, dm)
		conversationMap[partnerId] = conv
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// マップからスライスに変換
	var conversations []model.Conversation
	for _, conv := range conversationMap {
		conversations = append(conversations, conv)
	}

	return conversations, nil
}
