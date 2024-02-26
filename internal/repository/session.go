package session

import (
	"github.com/pkg/errors"
	"math/rand"
	"sync"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

var (
	ErrNoUser            = errors.New("no user")
	ErrUserAlreadyExists = errors.New("user exists")
	ErrNoSession         = errors.New("no session")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Handler struct {
	*sync.Mutex
	sessions map[string]uint
	users    *users
}

type users struct {
	storage map[string]models.User
	nextID  uint
}

func NewHandler() *Handler {
	return &Handler{
		sessions: make(map[string]uint, 10),
		users:    newUsers(),
	}
}

func newUsers() *users {
	return &users{
		storage: make(map[string]models.User),
	}
}

func (h *Handler) CreateUser(user models.User) (uint, error) {
	h.Lock()
	defer h.Unlock()
	if _, ok := h.users.storage[user.Username]; ok {
		return 0, ErrUserAlreadyExists
	}
	h.users.nextID++
	user.ID = h.users.nextID
	h.users.storage[user.Username] = user
	return user.ID, nil
}

func (h *Handler) GetUser(login string) (models.User, error) {
	h.Lock()
	user, ok := h.users.storage[login]
	h.Unlock()
	if !ok {
		return models.User{}, ErrNoUser
	}
	return user, nil
}

func (h *Handler) CreateSession(login string, userID uint) string {
	sID := RandStringRunes(32)
	h.Lock()
	h.sessions[sID] = userID
	h.Unlock()
	return sID
}

func (h *Handler) SessionExists(login string) bool {
	_, ok := h.sessions[login]
	return ok
}

func (h *Handler) DeleteSession(login string) error {
	h.Lock()
	defer h.Unlock()
	if _, ok := h.sessions[login]; !ok {
		return ErrNoSession
	}
	delete(h.sessions, login)
	return nil
}
