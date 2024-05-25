package dto

import (
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Notification struct {
	ID         uint      `json:"id"`
	Type       string    `json:"type"`
	ReadStatus string    `json:"readStatus"`
	Payload    string    `json:"payload"`
	CreatedAt  time.Time `json:"createdAt"`
}

func ConvertNotifications(notifications []models.Notification) []Notification {
	result := make([]Notification, 0, len(notifications))
	for _, notification := range notifications {
		result = append(result, Notification{
			ID:         notification.ID,
			Type:       notification.Type,
			ReadStatus: notification.ReadStatus,
			Payload:    notification.Payload,
			CreatedAt:  notification.CreatedAt,
		})
	}
	return result
}

type ReadNotificationInput struct {
	NotificationID uint `json:"notificationId"`
}
