package dao

import (
	"api/sqlc"
	"database/sql"
)

// New メソッドの実装
func (d *Dao) New(db sqlc.DBTX) *sqlc.Queries {
	return sqlc.New(d.db)
}

// WithTx メソッドの実装
func (d *Dao) WithTx(tx *sql.Tx) *sqlc.Queries {
	return d.queries.WithTx(tx)
}

// Begin メソッドの実装
func (d *Dao) Begin() (*sql.Tx, error) {
	return d.db.Begin()
}
