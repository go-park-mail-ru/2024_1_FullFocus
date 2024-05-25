package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type NotificationRepo struct {
	storage db.Database
}

func NewNotificationRepo(dbClient db.Database) *NotificationRepo {
	return &NotificationRepo{
		storage: dbClient,
	}
}

// SendNotification is TODO
// func (r *NotificationRepo) SendNotification(ctx context.Context, profileID uint, data []byte) error {}

func (r *NotificationRepo) GetAllNotifications(ctx context.Context, profileID uint) ([]models.Notification, error) {
	q := `SELECT n.id, n.type, n.read_status, n.payload, n.created_at
		  FROM notification n
		  WHERE profile_id = ?
		  ORDER BY created_at;`

	var notifications []dao.Notification
	if err := r.storage.Select(ctx, &notifications, q, profileID); err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrNoNotifications
	}
	if len(notifications) == 0 {
		return nil, models.ErrNoNotifications
	}
	return dao.ConvertNotifications(notifications), nil
}

func (r *NotificationRepo) MarkNotificationRead(ctx context.Context, id uint) error {
	q := `UPDATE notification n
		  SET read_status = true
		  WHERE n.id = ?;`

	if _, err := r.storage.Exec(ctx, q, id); err != nil {
		return models.ErrNoNotifications
	}
	return nil
}
