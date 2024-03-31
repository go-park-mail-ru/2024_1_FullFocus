package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	Close() error
	Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error)
	Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	GetRawDb() *sqlx.DB // tmp solution
}
