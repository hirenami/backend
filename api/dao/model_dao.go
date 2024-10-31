package dao

import (
	"api/sqlc"
	"database/sql"
)

type Dao struct {
	db *sql.DB
	queries *sqlc.Queries
}

// NewDao コンストラクタ
func NewDao(db *sql.DB, queries *sqlc.Queries) *Dao {
	return &Dao{
		db: db,
		queries: sqlc.New(db),
	}
}

