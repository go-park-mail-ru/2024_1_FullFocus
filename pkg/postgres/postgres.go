package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	_ "github.com/jackc/pgx" // postgres driver
	"github.com/jmoiron/sqlx"
)

type PgxDatabase struct {
	dsn    string
	client *sqlx.DB
}

func NewPgxDatabase(ctx context.Context, cfg config.PostgresConfig) (database.Database, error) {
	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		hostPort,
		cfg.Database,
		cfg.Sslmode,
	)
	dbClient, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, err
	}
	dbClient.SetMaxOpenConns(cfg.MaxOpenConns)
	dbClient.SetConnMaxIdleTime(time.Second * time.Duration(cfg.MaxIdleTime))
	return &PgxDatabase{
		dsn:    dsn,
		client: dbClient,
	}, nil
}

func (db *PgxDatabase) GetRawDB() *sqlx.DB {
	return db.client
}

func (db *PgxDatabase) Close() error {
	return db.client.Close()
}

func (db *PgxDatabase) Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error) {
	return db.client.ExecContext(ctx, sqlx.Rebind(sqlx.DOLLAR, q), args...)
}

func (db *PgxDatabase) Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	return db.client.GetContext(ctx, dest, sqlx.Rebind(sqlx.DOLLAR, q), args...)
}

func (db *PgxDatabase) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return db.client.NamedExecContext(ctx, query, arg)
}
