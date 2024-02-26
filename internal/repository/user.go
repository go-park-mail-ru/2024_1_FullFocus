package repository

import (
	"github.com/pkg/errors"
	"sync"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

var (
	ErrNoUser            = errors.New("no user")
	ErrUserAlreadyExists = errors.New("user exists")
)

type UserRepo struct {
	*sync.Mutex
	storage map[string]models.User
	nextID  uint
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		storage: make(map[string]models.User),
	}
}

func (r *UserRepo) CreateUser(user models.User) (uint, error) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.storage[user.Username]; ok {
		return 0, ErrUserAlreadyExists
	}
	r.nextID++
	user.ID = r.nextID
	r.storage[user.Username] = user
	return user.ID, nil
}

func (r *UserRepo) GetUser(login string) (models.User, error) {
	r.Lock()
	user, ok := r.storage[login]
	r.Unlock()
	if !ok {
		return models.User{}, ErrNoUser
	}
	return user, nil
}
