package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

func (r *AuthRepo) CreateSession(ctx context.Context, uID uint) string {
	sID := uuid.New().String()
	start := time.Now()
	r.redis.Set(sID, uID, r.sessionTTL)
	logger.Info(ctx, fmt.Sprintf("session inserted in %s", time.Since(start)))
	return sID
}

func (r *AuthRepo) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	start := time.Now()
	uID, err := r.redis.Get(sID).Uint64()
	if err != nil {
		logger.Error(ctx, "no session found")
		return 0, models.ErrNoSession
	}
	logger.Info(ctx, fmt.Sprintf("user_id selected in %s", time.Since(start)))
	return uint(uID), nil
}

func (r *AuthRepo) SessionExists(ctx context.Context, sID string) bool {
	start := time.Now()
	_, err := r.redis.Get(sID).Uint64()
	logger.Info(ctx, fmt.Sprintf("session checked in %s", time.Since(start)))
	if err != nil {
		logger.Info(ctx, "no session")
		return false
	}
	logger.Info(ctx, "session found")
	return true
}

func (r *AuthRepo) DeleteSession(ctx context.Context, sID string) error {
	start := time.Now()
	if err := r.redis.Get(sID).Err(); err != nil {
		logger.Error(ctx, "no session found")
		return models.ErrNoSession
	}
	logger.Info(ctx, fmt.Sprintf("session checked in %s", time.Since(start)))
	start = time.Now()
	r.redis.Del(sID)
	logger.Info(ctx, fmt.Sprintf("session deleted in %s", time.Since(start)))

	return nil
}
