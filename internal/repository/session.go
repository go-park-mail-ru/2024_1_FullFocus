package repository

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"sync"

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

func (r *SessionRepo) CreateSession(userID uint) string {
	sID := uuid.New().String()
	r.Lock()
	r.sessions[sID] = userID
	r.Unlock()
	return sID
}

func (r *SessionRepo) SessionExists(sID string) bool {
	r.Lock()
	_, ok := r.sessions[sID]
	r.Unlock()
	return ok
}

func (r *SessionRepo) DeleteSession(sID string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.sessions[sID]; !ok {
		return models.ErrNoSession
	}
	delete(r.sessions, sID)
	return nil
}
