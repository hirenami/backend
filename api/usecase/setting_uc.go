package usecase

import (
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateAccount(ctx context.Context, firebaseUid string, username string, userId string) error {
	//バリデーション
	if username == "" {
		return errors.New("username is required")
	}
	if userId == "" {
		return errors.New("userId is required")
	}

	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//Daoのメソッドを呼び出し
	err = u.dao.CreateAccount(ctx, tx, firebaseUid, username, userId)
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

// Usecase メソッドの実装
func (u *Usecase) GetIdByUID(ctx context.Context, firebaseUid string) (string, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return "", err
	}

	//Daoのメソッドを呼び出し
	userId, err := u.dao.GetIdbyUid(ctx, tx, firebaseUid)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return "", err
	}
	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return userId, nil
}

// Usecase メソッドの実装
func (u *Usecase) DeleteAccount(ctx context.Context, myId string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//Daoのメソッドを呼び出し
	err = u.dao.CreateIsDeleted(ctx, tx, true, myId)
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

// Usecase メソッドの実装
func (u *Usecase) UpdatePrivateUsecase (ctx context.Context, myId string, isPrivate bool) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//Daoのメソッドを呼び出し
	err = u.dao.CreateIsPrivate(ctx, tx, isPrivate, myId)
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
