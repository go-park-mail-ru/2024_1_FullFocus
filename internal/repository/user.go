package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/google/uuid"
)

type UserRepo struct {
	storage db.Database
}

func NewUserRepo(dbClient db.Database) *UserRepo {
	return &UserRepo{
		storage: dbClient,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	l := helper.GetLoggerFromContext(ctx)
	userRow := ConvertUserToTable(user)
	q := `INSERT INTO default_user (id, user_login, password_hash) VALUES ($1, $2, $3);`
	l.Info(q, slog.String("args", fmt.Sprintf("$1 = %s $2 = %s, $3 = %s", userRow.Id, userRow.Login, userRow.Password_hash)))
	start := time.Now()
	defer func() {
		l.Info(fmt.Sprintf("created in %s", time.Since(start)))
	}()
	_, err := r.storage.Exec(ctx, q, userRow)
	if err != nil {
		l.Error("user already exists")
		return uuid.Nil, models.ErrUserAlreadyExists
	}
	return userRow.Id, nil
}

func (r *UserRepo) GetUser(ctx context.Context, username string) (models.User, error) {
	l := helper.GetLoggerFromContext(ctx)
	q := `SELECT id FROM default_user WHERE user_login = $1;`
	l.Info(q, slog.String("args", fmt.Sprintf("$1 = %s", username)))
	start := time.Now()
	defer func() {
		l.Info(fmt.Sprintf("queried in %s", time.Since(start)))
	}()
	userRow := &UserTable{}
	err := r.storage.Get(ctx, userRow, q, username)
	if err != nil {
		l.Error("user not found")
		return models.User{}, models.ErrNoUser
	}
	return ConvertTableToUser(*userRow), nil
}
