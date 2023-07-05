package db

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries { return &Queries{db} }

type Queries struct{ db DBTX }

func (q *Queries) WithTx(tx *sql.Tx) *Queries { return &Queries{tx} }