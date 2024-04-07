package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=database.go -destination=./mocks/database_mock.go
type Database interface {
	Close() error
	Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error)
	Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	GetRawDB() *sqlx.DB // tmp solution
}
