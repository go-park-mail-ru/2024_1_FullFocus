package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"log/slog"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
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
	l := helper.GetLoggerFromContext(ctx)
	l.Info(`INSERT INTO user (username, password) VALUES ($1, $2);`,
		slog.String("args", fmt.Sprintf("$1 = %s, $2 = %s", user.Username, user.Password)))

	start := time.Now()
	defer func() {
		l.Info(fmt.Sprintf("created in %s", time.Since(start)))
	}()
	r.Lock()
	defer r.Unlock()
	if _, ok := r.storage[user.Username]; ok {
		l.Error("user already exists")
		return 0, models.ErrUserAlreadyExists
	}
	r.nextID++
	r.storage[user.Username] = user
	return user.ID, nil
}

func (r *UserRepo) GetUser(ctx context.Context, username string) (models.User, error) {
	l := helper.GetLoggerFromContext(ctx)
	l.Info(`SELECT * FROM user WHERE usename = $1;`,
		slog.String("args", fmt.Sprintf("$1 = %s", username)))

	start := time.Now()
	r.Lock()
	user, ok := r.storage[username]
	r.Unlock()
	if !ok {
		l.Error("user not found")
		return models.User{}, models.ErrNoUser
	}
	l.Info(fmt.Sprintf("queried in %s", time.Since(start)))
	return user, nil
}
