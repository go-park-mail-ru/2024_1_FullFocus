package helper

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type UserID struct{}

func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(UserID{}).(uint)
	if !ok {
		return 0, models.ErrNoUserID
	}
	return userID, nil
}
