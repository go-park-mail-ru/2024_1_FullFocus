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

// Session godoc
// @Tags Session
// @Summary Make new session Rep
// @Success 200 {object} map[string]uint
// @Router /NewSessionRepo [post]
func NewSessionRepo() *SessionRepo {
	return &SessionRepo{
		sessions: make(map[string]uint, 10),
	}
}

// Session godoc
// @Tags Session
// @Summary Make new session
// @Param userID body uint true "UserID"
// @Success 200 {object} string
// @Router /CreateSession [post]
func (r *SessionRepo) CreateSession(userID uint) string {
	sID := uuid.New().String()
	r.Lock()
	r.sessions[sID] = userID
	r.Unlock()
	return sID
}

// Session godoc
// @Tags Session
// @Summary Checking for a created session
// @Param sID body string true "sID"
// @Success 200 {object} bool
// @Router /SessionExists [post]
func (r *SessionRepo) SessionExists(sID string) bool {
	r.Lock()
	_, ok := r.sessions[sID]
	r.Unlock()
	return ok
}

// Session godoc
// @Tags Session
// @Summary Delete session
// @Param sID body string true "sID"
// @Success 200
// @Router /DeleteSession [post]
func (r *SessionRepo) DeleteSession(sID string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.sessions[sID]; !ok {
		return models.ErrNoSession
	}
	delete(r.sessions, sID)
	return nil
}
