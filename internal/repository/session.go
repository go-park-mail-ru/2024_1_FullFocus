package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

type SessionRepo struct {
	client     *redis.Client
	sessionTTL time.Duration
}

func NewSessionRepo(c *redis.Client, sessTTL time.Duration) *SessionRepo {
	return &SessionRepo{
		client:     c,
		sessionTTL: sessTTL,
	}
}

func (r *SessionRepo) CreateSession(ctx context.Context, userID uuid.UUID) string {
	sID := uuid.New().String()
	start := time.Now()
	r.client.Set(sID, userID, r.sessionTTL)
	logger.Info(ctx, fmt.Sprintf("session inserted in %s", time.Since(start)))
	return sID
}

func (r *SessionRepo) GetUserIDBySessionID(ctx context.Context, sID string) (uuid.UUID, error) {
	start := time.Now()
	uIDstr := r.client.Get(sID).Val()
	logger.Info(ctx, fmt.Sprintf("user_id selected in %s", time.Since(start)))
	if uIDstr == "" {
		logger.Error(ctx, "no session found")
		return uuid.Nil, models.ErrNoSession
	}
	uID, _ := uuid.Parse(uIDstr)
	return uID, nil
}

func (r *SessionRepo) SessionExists(ctx context.Context, sID string) bool {
	start := time.Now()
	_, err := r.client.Get(sID).Uint64()
	logger.Info(ctx, fmt.Sprintf("session checked in %s", time.Since(start)))
	if err != nil {
		logger.Info(ctx, "no session")
		return false
	}
	logger.Info(ctx, "session found")
	return true
}

func (r *SessionRepo) DeleteSession(ctx context.Context, sID string) error {
	start := time.Now()
	if err := r.client.Get(sID).Err(); err != nil {
		logger.Error(ctx, "no session found")
		return models.ErrNoSession
	}
	logger.Info(ctx, fmt.Sprintf("session checked in %s", time.Since(start)))
	start = time.Now()
	r.client.Del(sID)
	logger.Info(ctx, fmt.Sprintf("session deleted in %s", time.Since(start)))

	return nil
}
