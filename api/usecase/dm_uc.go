package usecase

import (
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
	dm, err := u.dao.GetDms(ctx, tx, userId, repid)
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

	return dm, nil
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
