package helper

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	middleware "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware/auth"
)

func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(middleware.ContextUserKey).(uint)
	if !ok {
		return 0, models.ErrNoUserID
	}
	return userID, nil
}
