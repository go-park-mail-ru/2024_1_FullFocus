package repository

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

type UserRepo struct {
	nextID uint
	sync.Mutex
	storage map[string]models.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		storage: make(map[string]models.User),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user models.User) (uint, error) {
	logger.Info(ctx, `INSERT INTO user (username, password) VALUES ($1, $2);`,
		slog.String("args", fmt.Sprintf("$1 = %s, $2 = %s", user.Username, user.Password)))

	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("created in %s", time.Since(start)))
	}()
	r.Lock()
	defer r.Unlock()
	if _, ok := r.storage[user.Username]; ok {
		logger.Error(ctx, "user already exists")
		return 0, models.ErrUserAlreadyExists
	}
	r.nextID++
	r.storage[user.Username] = user
	return user.ID, nil
}

func (r *UserRepo) GetUser(ctx context.Context, username string) (models.User, error) {
	logger.Info(ctx, `SELECT * FROM user WHERE usename = $1;`,
		slog.String("args", "$1 = "+username))

	start := time.Now()
	r.Lock()
	user, ok := r.storage[username]
	r.Unlock()
	if !ok {
		logger.Error(ctx, "user not found")
		return models.User{}, models.ErrNoUser
	}
	logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	return user, nil
}
