package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PgxDatabase struct {
	dsn    string
	client *pgxpool.Pool
}

func GetDSN(cfg config.PostgresConfig) string {
	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s&search_path=%s",
		cfg.User,
		cfg.Password,
		hostPort,
		cfg.Database,
		cfg.Sslmode,
		cfg.SearchPath,
	)
}

func NewPgxDatabase(ctx context.Context, dbCfg *pgxpool.Config) (database.Database, error) {
	dbClient, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, err
	}
	return &PgxDatabase{
		dsn:    dbClient.Config().ConnString(),
		client: dbClient,
	}, nil
}

func (db *PgxDatabase) GetRawDB() *pgxpool.Pool {
	return db.client
}

func (db *PgxDatabase) Close() error {
	return db.client.Close()
}

func (db *PgxDatabase) Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error) {
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("%v", args)))
	start := time.Now()
	sqlRes, err := db.client.ExecContext(ctx, sqlx.Rebind(sqlx.DOLLAR, q), args...)
	logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	return sqlRes, err
}

func (db *PgxDatabase) Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("%v", args)))
	start := time.Now()
	err := db.client.GetContext(ctx, dest, sqlx.Rebind(sqlx.DOLLAR, q), args...)
	logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	return err
}

func (db *PgxDatabase) Select(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("%v", args)))
	start := time.Now()
	err := db.client.SelectContext(ctx, dest, sqlx.Rebind(sqlx.DOLLAR, q), args...)
	logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	return err
}

func (db *PgxDatabase) NamedExec(ctx context.Context, q string, arg interface{}) (sql.Result, error) {
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("%v", arg)))
	start := time.Now()
	result, err := db.client.NamedExecContext(ctx, q, arg)
	logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	return result, err
}

func (db *PgxDatabase) Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return db.client.BeginTxx(ctx, opts)
}
