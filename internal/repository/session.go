package repository

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/uuidgen"
	"github.com/pkg/errors"
	"sync"
)

var (
	ErrNoSession = errors.New("no session")
)

type SessionRepo struct {
	*sync.Mutex
	sessions map[string]uint
}

func NewSessionRepo() *SessionRepo {
	return &SessionRepo{
		sessions: make(map[string]uint, 10),
	}
}

func (r *SessionRepo) CreateSession(login string, userID uint) string {
	sID := uuidgen.RandStringRunes(32)
	r.Lock()
	r.sessions[sID] = userID
	r.Unlock()
	return sID
}

func (r *SessionRepo) SessionExists(login string) bool {
	_, ok := r.sessions[login]
	return ok
}

func (r *SessionRepo) DeleteSession(login string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.sessions[login]; !ok {
		return ErrNoSession
	}
	delete(r.sessions, login)
	return nil
}
