package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
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

func (r *SessionRepo) CreateSession(ctx context.Context, userID uint) string {
	l := helper.GetLoggerFromContext(ctx)
	sID := uuid.New().String()
	start := time.Now()
	r.client.Set(sID, userID, r.sessionTTL)
	l.Info(fmt.Sprintf("session inserted in %s", time.Since(start)))
	return sID
}

func (r *SessionRepo) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	l := helper.GetLoggerFromContext(ctx)
	start := time.Now()
	uID, err := r.client.Get(sID).Uint64()
	l.Info(fmt.Sprintf("user_id selected in %s", time.Since(start)))
	if err != nil {
		l.Error("no session found")
		return 0, models.ErrNoSession
	}
	return uint(uID), nil
}

func (r *SessionRepo) SessionExists(ctx context.Context, sID string) bool {
	l := helper.GetLoggerFromContext(ctx)
	start := time.Now()
	_, err := r.client.Get(sID).Uint64()
	l.Info(fmt.Sprintf("session checked in %s", time.Since(start)))
	if err != nil {
		l.Info("session found")
		return false
	}
	l.Info("no session")
	return true
}

func (r *SessionRepo) DeleteSession(ctx context.Context, sID string) error {
	l := helper.GetLoggerFromContext(ctx)
	start := time.Now()
	if err := r.client.Get(sID).Err(); err != nil {
		l.Error("no session found")
		return models.ErrNoSession
	}
	l.Info(fmt.Sprintf("session checked in %s", time.Since(start)))
	start = time.Now()
	r.client.Del(sID)
	l.Info(fmt.Sprintf("session deleted in %s", time.Since(start)))

	return nil
}
