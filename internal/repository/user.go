package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"sync"

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
	l := logger.LoggerFromContext(ctx)
	r.Lock()
	defer r.Unlock()
	if _, ok := r.storage[user.Username]; ok {
		l.Error("user already exists")
		return 0, models.ErrUserAlreadyExists
	}
	r.nextID++
	r.storage[user.Username] = user
	l.Info(fmt.Sprintf("user created: %d", user.ID))
	return user.ID, nil
}

func (r *UserRepo) GetUser(ctx context.Context, username string) (models.User, error) {
	l := logger.LoggerFromContext(ctx)
	r.Lock()
	user, ok := r.storage[username]
	r.Unlock()
	if !ok {
		l.Error("user not found")
		return models.User{}, models.ErrNoUser
	}
	l.Info(fmt.Sprintf("user found: %d", user.ID))
	return user, nil
}
