package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrNoSession = errors.New("no session")
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
	_, ok := r.sessions[sID]
	return ok
}

func (r *SessionRepo) DeleteSession(sID string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.sessions[sID]; !ok {
		return ErrNoSession
	}
	delete(r.sessions, sID)
	return nil
}
