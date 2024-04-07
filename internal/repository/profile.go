package repository

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"sync"
)

type ProfileRepo struct {
	nextID uint
	sync.Mutex
	storage map[string]models.Profile
}

func NewProfileRepo() *ProfileRepo {
	return &ProfileRepo{
		storage: make(map[string]models.Profile),
	}
}

func (r *ProfileRepo) CreateProfile(ctx context.Context, profile models.Profile) (uint, error) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.storage[profile.User.Username]; ok {
		logger.Error(ctx, "profile already exists")
		return 0, models.ErrUserAlreadyExists
	}
	r.nextID++
	r.storage[profile.User.Username] = profile
	return profile.User.ID, nil
}

func (r *ProfileRepo) GetProfile(ctx context.Context, username string) (models.Profile, error) {
	r.Lock()
	defer r.Unlock()
	profile, ok := r.storage[username]
	if !ok {
		logger.Error(ctx, "user not found")
		return models.Profile{}, models.ErrNoProfile
	}
	return profile, nil
}

func (r *ProfileRepo) UpdateProfile(ctx context.Context, username string, profileNew models.Profile) error {
	r.Lock()
	defer r.Unlock()
	_, ok := r.storage[username]
	if !ok {
		logger.Error(ctx, "user not found")
		return models.ErrNoProfile
	}
	r.storage[username] = profileNew
	return nil
}
