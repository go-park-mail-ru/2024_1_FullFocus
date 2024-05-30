package repository

import (
	"time"

	"github.com/go-redis/redis"

	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
)

type AuthRepo struct {
	redis      *redis.Client
	storage    db.Database
	sessionTTL time.Duration
}

func NewAuthRepo(rc *redis.Client, s db.Database, sessTTL time.Duration) *AuthRepo {
	return &AuthRepo{
		redis:      rc,
		storage:    s,
		sessionTTL: sessTTL,
	}
}
