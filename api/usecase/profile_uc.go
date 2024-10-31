package usecase

import (
	"api/sqlc"
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateProfileUsecase(ctx context.Context, userId, username, biography, header_url, icon_url string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
	}

	//バリデーション
	if username == "" {
		return errors.New("username is required")
	}
	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("user does not exist")
	}

	//Daoのメソッドを呼び出し
	err = u.dao.CreateBiography(ctx, tx, userId, biography)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.CreateHeaderImage(ctx, tx, userId, header_url)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.CreateIconImage(ctx, tx, userId, icon_url)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.UpdateUsername(ctx, tx, username, userId)
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

func (u *Usecase) GetProfileUsecase(ctx context.Context, userId string) (sqlc.User, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return sqlc.User{}, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return sqlc.User{}, err
	} else if !bool {
		return sqlc.User{}, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	user, err := u.dao.GetProfile(ctx, tx, userId)
	if err != nil {
		return sqlc.User{}, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
