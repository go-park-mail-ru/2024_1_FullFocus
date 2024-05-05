package repository_test

// import (
// 	"context"
// 	"math/rand"
// 	"testing"
// 	"time"

// 	"github.com/alicebob/miniredis/v2"
// 	"github.com/go-redis/redis"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"

// 	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/models"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/repository"
// )

// const _sessionTTL = 24 * time.Hour

// func TestNewSessionRepo(t *testing.T) {
// 	mr := miniredis.RunT(t)
// 	t.Cleanup(func() {
// 		mr.Close()
// 	})
// 	t.Run("Check SessionRepo creation", func(t *testing.T) {
// 		rc := redis.NewClient(&redis.Options{
// 			Addr: mr.Addr(),
// 		})
// 		defer func() {
// 			_ = rc.Close()
// 		}()
// 		sr := repository.NewSessionRepo(rc, _sessionTTL)
// 		require.NotEmpty(t, sr, "sessionrepo not created")
// 	})
// }

// func TestCreateSession(t *testing.T) {
// 	mr := miniredis.RunT(t)
// 	t.Cleanup(func() {
// 		mr.Close()
// 	})
// 	t.Run("Check valid sessionID by random userID", func(t *testing.T) {
// 		rc := redis.NewClient(&redis.Options{
// 			Addr: mr.Addr(),
// 		})
// 		defer func() {
// 			_ = rc.Close()
// 		}()
// 		sID := repository.NewSessionRepo(rc, _sessionTTL).CreateSession(context.Background(), uint(rand.Uint32()))
// 		_, err := uuid.Parse(sID)
// 		require.NoError(t, err, "got an empty sessionID")
// 	})
// }

// func TestSessionExists(t *testing.T) {
// 	mr := miniredis.RunT(t)
// 	t.Cleanup(func() {
// 		mr.Close()
// 	})
// 	rc := redis.NewClient(&redis.Options{
// 		Addr: mr.Addr(),
// 	})
// 	defer func() {
// 		_ = rc.Close()
// 	}()
// 	sr := repository.NewSessionRepo(rc, _sessionTTL)
// 	t.Run("Check real sessionID in SessionRepo", func(t *testing.T) {
// 		sID := sr.CreateSession(context.Background(), uint(rand.Uint32()))
// 		got := sr.SessionExists(context.Background(), sID)
// 		require.True(t, got, "valid session not found")
// 	})
// 	t.Run("Check empty sessionID in SessionRepo", func(t *testing.T) {
// 		got := sr.SessionExists(context.Background(), "")
// 		require.False(t, got, "found empty session")
// 	})
// }

// func TestDeleteSession(t *testing.T) {
// 	mr := miniredis.RunT(t)
// 	t.Cleanup(func() {
// 		mr.Close()
// 	})
// 	rc := redis.NewClient(&redis.Options{
// 		Addr: mr.Addr(),
// 	})
// 	defer func() {
// 		_ = rc.Close()
// 	}()
// 	sr := repository.NewSessionRepo(rc, _sessionTTL)
// 	t.Run("Check existing sessionID delete", func(t *testing.T) {
// 		sID := sr.CreateSession(context.Background(), uint(rand.Uint32()))
// 		err := sr.DeleteSession(context.Background(), sID)
// 		require.NoError(t, err, "existing sID not deleted")
// 	})
// 	t.Run("Check empty sessionID delete", func(t *testing.T) {
// 		err := sr.DeleteSession(context.Background(), "")
// 		require.Equal(t, models.ErrNoSession, err, "found empty sID")
// 	})
// }
