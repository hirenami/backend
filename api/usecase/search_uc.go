package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) SearchByKeywordUsecase(ctx context.Context, keyword string) ([]sqlc.Tweet,error) {

	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil,err
	}

	if keyword == "" {
		return nil,errors.New("keyword is required")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweets,err := u.dao.SearchByKeyword(ctx, tx, keyword)
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

	return tweets,nil
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
func (u *Usecase) SearchByHashtagUsecase(ctx context.Context, keyword string) ([]sqlc.Tweet,error) {
	
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil,err
	}

	if keyword == "" {
		return nil,errors.New("keyword is required")
	}

	// daoのメソッドにトランザクションを渡して実行
	tweets,err := u.dao.SearchByHashtag(ctx, tx, keyword)
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

	return tweets,nil
}