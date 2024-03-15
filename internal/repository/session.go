package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"sync"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/google/uuid"
)

type SessionRepo struct {
	sync.Mutex
	sessions map[string]uint
}

func NewSessionRepo() *SessionRepo {
	return &SessionRepo{
		sessions: make(map[string]uint, 10),
	}
}

func (r *SessionRepo) CreateSession(ctx context.Context, userID uint) string {
	l := helper.GetLoggerFromContext(ctx)
	sID := uuid.New().String()
	r.Lock()
	r.sessions[sID] = userID
	r.Unlock()
	l.Info(fmt.Sprintf("session created: %s", sID))
	return sID
}

func (r *SessionRepo) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	l := helper.GetLoggerFromContext(ctx)
	r.Lock()
	defer r.Unlock()
	uID, ok := r.sessions[sID]
	if !ok {
		l.Error("no session")
		return 0, models.ErrNoSession
	}
	l.Info(fmt.Sprintf("user found: %d", uID))
	return uID, nil
}

func (r *SessionRepo) SessionExists(ctx context.Context, sID string) bool {
	l := helper.GetLoggerFromContext(ctx)
	r.Lock()
	_, ok := r.sessions[sID]
	r.Unlock()
	l.Info(fmt.Sprintf("session found: %T", ok))
	return ok
}

func (r *SessionRepo) DeleteSession(ctx context.Context, sID string) error {
	l := helper.GetLoggerFromContext(ctx)
	r.Lock()
	defer r.Unlock()
	if _, ok := r.sessions[sID]; !ok {
		l.Error("no session found")
		return models.ErrNoSession
	}
	delete(r.sessions, sID)
	l.Info("session deleted")
	return nil
}
