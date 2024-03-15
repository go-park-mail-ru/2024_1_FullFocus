package repository

import (
	"context"
	"math/rand"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewSessionRepo(t *testing.T) {
	t.Run("Check SessionRepo creation", func(t *testing.T) {
		sr := NewSessionRepo()
		require.NotEmpty(t, sr, "sessionrepo not created")
	})
}

func TestCreateSession(t *testing.T) {
	t.Run("Check valid sessionID by random userID", func(t *testing.T) {
		sID := NewSessionRepo().CreateSession(context.Background(), uint(rand.Uint32()))
		_, err := uuid.Parse(sID)
		require.Equal(t, nil, err, "got an empty sessionID")
	})
}

func TestSessionExists(t *testing.T) {
	sr := NewSessionRepo()
	t.Run("Check real sessionID in SessionRepo", func(t *testing.T) {
		sID := sr.CreateSession(context.Background(), uint(rand.Uint32()))
		got := sr.SessionExists(context.Background(), sID)
		require.Equal(t, true, got, "valid session not found")
	})
	t.Run("Check empty sessionID in SessionRepo", func(t *testing.T) {
		got := sr.SessionExists(context.Background(), "")
		require.Equal(t, false, got, "found empty session")
	})
}

func TestDeleteSession(t *testing.T) {
	sr := NewSessionRepo()
	t.Run("Check existing sessionID delete", func(t *testing.T) {
		sID := sr.CreateSession(context.Background(), uint(rand.Uint32()))
		err := sr.DeleteSession(context.Background(), sID)
		require.Equal(t, nil, err, "existing sID not deleted")
	})
	t.Run("Check empty sessionID delete", func(t *testing.T) {
		err := sr.DeleteSession(context.Background(), "")
		require.Equal(t, models.ErrNoSession, err, "found empty sID")
	})
}
