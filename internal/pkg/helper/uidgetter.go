package helper

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(models.ContextUserKey).(uint)
	if !ok {
		return 0, models.ErrNoUserID
	}
	return userID, nil
}
