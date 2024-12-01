package usecase

import (
	"context"
	"errors"
	"log"
)

// Usecase メソッドの実装
func (u *Usecase) CreateBlockUsecase(ctx context.Context, userId string, blockId string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
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

	if bool, err := u.dao.IsUserExists(ctx, tx, blockId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("block user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.CreateBlock(ctx, tx, userId, blockId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.DeleteFollow(ctx, tx, userId, blockId)
	if err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	}

	err = u.dao.DeleteFollow(ctx, tx, blockId, userId)
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

func (u *Usecase) DeleteBlockUsecase(ctx context.Context, userId string, blockId string) error {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return err
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

	if bool, err := u.dao.IsUserExists(ctx, tx, blockId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return err
	} else if !bool {
		return errors.New("block user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	err = u.dao.DeleteBlock(ctx, tx, userId, blockId)
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

func (u *Usecase) IsBlockedckUsecase(ctx context.Context, userId, blockid string) (bool, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return false, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return false, err
	} else if !bool {
		return false, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	block, err := u.dao.IsBlocked(ctx, tx, userId, blockid)
	if err != nil {
		return false, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return block, nil
}

func (u *Usecase) GetBlocksUsecase(ctx context.Context, userId string) ([]string, error) {
	// トランザクションを開始
	tx, err := u.dao.Begin()
	if err != nil {
		return nil, err
	}

	if bool, err := u.dao.IsUserExists(ctx, tx, userId); err != nil {
		// エラーが発生した場合、ロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("ロールバック中にエラーが発生しました: %v", rbErr)
		}
		return nil, err
	} else if !bool {
		return nil, errors.New("user does not exist")
	}

	// daoのメソッドにトランザクションを渡して実行
	blocks, err := u.dao.GetBlocks(ctx, tx, userId)
	if err != nil {
		return nil, err
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return blocks, nil
}
