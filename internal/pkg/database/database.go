//go:generate mockgen -source=database.go -destination=./mocks/database_mock.go
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
	Select(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	GetRawDB() *sqlx.DB // tmp solution
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}
