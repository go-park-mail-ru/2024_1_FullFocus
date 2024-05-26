package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type NotificationUsecase struct {
	notificationRepo repository.Notifications
}

func NewNotificationUsecase(nr repository.Notifications) *NotificationUsecase {
	return &NotificationUsecase{
		notificationRepo: nr,
	}
}

func (u *NotificationUsecase) GetAllNotifications(ctx context.Context, profileID uint) ([]models.Notification, error) {
	return u.notificationRepo.GetAllNotifications(ctx, profileID)
}

func (u *NotificationUsecase) MarkNotificationRead(ctx context.Context, id uint) error {
	return u.notificationRepo.MarkNotificationRead(ctx, id)
}
