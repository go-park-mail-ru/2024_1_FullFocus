package repository

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
)

func TestNewSessionRepo(t *testing.T) {
	t.Run("Check SessionRepo creation", func(t *testing.T) {
		sr := NewSessionRepo()
		if sr == nil {
			t.Errorf("Sessionrepo not created")
		}
	})
}

func TestCreateSession(t *testing.T) {
	t.Run("Check valid sessionID by random userID", func(t *testing.T) {
		sID := NewSessionRepo().CreateSession(uint(rand.Uint32()))
		if _, err := uuid.Parse(sID); err != nil {
			t.Errorf("Got an empty sessionID")
		}
	})
}

func TestSessionExists(t *testing.T) {
	sr := NewSessionRepo()
	t.Run("Check real sessionID in SessionRepo", func(t *testing.T) {
		sID := sr.CreateSession(uint(rand.Uint32()))
		if !sr.SessionExists(sID) {
			t.Errorf("Got incorrect result")
		}
	})
	t.Run("Check empty sessionID in SessionRepo", func(t *testing.T) {
		sID := ""
		if sr.SessionExists(sID) {
			t.Errorf("Got incorrect result")
		}
	})
}

func TestDeleteSession(t *testing.T) {
	sr := NewSessionRepo()
	t.Run("Check existing sessionID delete", func(t *testing.T) {
		sID := sr.CreateSession(uint(rand.Uint32()))
		if err := sr.DeleteSession(sID); err != nil {
			t.Errorf("existing sID not deleted")
		}
	})
	t.Run("Check empty sessionID delete", func(t *testing.T) {
		if err := sr.DeleteSession(""); err != ErrNoSession {
			t.Errorf("found empty sID")
		}
	})
}
