package usecase

import (
	"api/dao"
)

// Usecase構造体
type Usecase struct {
	dao *dao.Dao
}

// NewTestUsecase コンストラクタ
func NewUsecase(dao *dao.Dao) *Usecase {
	return &Usecase{
		dao: dao,
	}
}