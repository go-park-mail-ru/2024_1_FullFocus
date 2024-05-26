package dao

import (
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Notification struct {
	ID         uint      `db:"id"`
	Type       string    `db:"type"`
	ReadStatus bool      `db:"read_status"`
	Payload    string    `db:"payload"`
	CreatedAt  time.Time `db:"created_at"`
}

func ConvertNotifications(notifications []Notification) []models.Notification {
	result := make([]models.Notification, 0, len(notifications))
	for _, notification := range notifications {
		result = append(result, models.Notification{
			ID:         notification.ID,
			Type:       notification.Type,
			ReadStatus: notification.ReadStatus,
			Payload:    notification.Payload,
			CreatedAt:  notification.CreatedAt,
		})
	}
	return result
}
