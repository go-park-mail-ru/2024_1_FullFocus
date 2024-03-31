package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type PgxDatabase struct {
	dsn    string
	client *sqlx.DB
}

func NewPgxDatabase(ctx context.Context, cfg config.PostgresConfig) (database.Database, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Sslmode,
	)
	dbClient, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, err
	}
	dbClient.SetMaxOpenConns(10)
	dbClient.SetConnMaxIdleTime(10 * time.Second)
	return &PgxDatabase{
		dsn:    dsn,
		client: dbClient,
	}, nil
}

func (db *PgxDatabase) GetRawDb() *sqlx.DB {
	return db.client
}

func (db *PgxDatabase) Close() error {
	return db.client.Close()
}

func (db *PgxDatabase) Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error) {
	return db.client.ExecContext(ctx, sqlx.Rebind(sqlx.DOLLAR, q), args)
}

func (db *PgxDatabase) Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	return db.client.GetContext(ctx, dest, sqlx.Rebind(sqlx.DOLLAR, q), args)
}
