package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"log/slog"
	"sync"
	"time"

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
	l.Info(`INSERT INTO session (sess_id, user_id) VALUES ($1, $2);`,
		slog.String("args", fmt.Sprintf("$1 = %s, $2 = %d", sID, userID)))

	start := time.Now()
	r.Lock()
	r.sessions[sID] = userID
	r.Unlock()
	l.Info(fmt.Sprintf("session inserted in %s", time.Since(start)))

	return sID
}

func (r *SessionRepo) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	l := helper.GetLoggerFromContext(ctx)
	l.Info(`SELECT user_id FROM session WHERE sess_id = $1;`,
		slog.String("args", fmt.Sprintf("$1 = %s", sID)))

	start := time.Now()
	r.Lock()
	uID, ok := r.sessions[sID]
	r.Unlock()
	l.Info(fmt.Sprintf("user_id selected in %s", time.Since(start)))

	if !ok {
		l.Error("no session found")
		return 0, models.ErrNoSession
	}
	return uID, nil
}

func (r *SessionRepo) SessionExists(ctx context.Context, sID string) bool {
	l := helper.GetLoggerFromContext(ctx)
	l.Info(`SELECT COUNT(user_id) FROM session WHERE sess_id = $1;`,
		slog.String("args", fmt.Sprintf("$1 = %s", sID)))

	start := time.Now()
	r.Lock()
	_, ok := r.sessions[sID]
	r.Unlock()
	l.Info(fmt.Sprintf("session checked in %s", time.Since(start)))
	return ok
}

func (r *SessionRepo) DeleteSession(ctx context.Context, sID string) error {
	l := helper.GetLoggerFromContext(ctx)
	l.Info(`DELETE FROM session WHERE sess_id = $1;`,
		slog.String("args", fmt.Sprintf("$1 = %s", sID)))

	start := time.Now()
	defer func() {
		l.Info(fmt.Sprintf("session deleted in %s", time.Since(start)))
	}()
	r.Lock()
	defer r.Unlock()
	if _, ok := r.sessions[sID]; !ok {
		l.Error("no session found")
		return models.ErrNoSession
	}
	delete(r.sessions, sID)

	return nil
}
