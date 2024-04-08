package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
)

type UserRepo struct {
	storage db.Database
}

func NewUserRepo(dbClient db.Database) *UserRepo {
	return &UserRepo{
		storage: dbClient,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user models.User) (uint, error) {
	userRow := dao.ConvertUserToTable(user)
	q := `INSERT INTO default_user (user_login, password_hash) VALUES ($1, $2);`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %s", userRow.Login)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("created in %s", time.Since(start)))
	}()
	_, err := r.storage.Exec(ctx, q, userRow.Login, userRow.PasswordHash)
	if err != nil {
		logger.Error(ctx, "user already exists")
		return 0, models.ErrUserAlreadyExists
	}
	return userRow.ID, nil
}

func (r *UserRepo) GetUser(ctx context.Context, username string) (models.User, error) {
	q := `SELECT id, password_hash FROM default_user WHERE user_login = $1;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %s", username)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	}()
	userRow := &dao.UserTable{}
	if err := r.storage.Get(ctx, userRow, q, username); err != nil {
		logger.Error(ctx, "user not found")
		return models.User{}, models.ErrNoUser
	}
	return dao.ConvertTableToUser(*userRow), nil
}
