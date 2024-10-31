package usecase

import (
	"context"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) Test(ctx context.Context, userId, content, media_url string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.Test(ctx, tx, userId, content, media_url)
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
