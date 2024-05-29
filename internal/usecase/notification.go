package usecase

import (
	"context"
	"fmt"
	"strconv"

	centrifuge "github.com/centrifugal/gocent/v3"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type NotificationUsecase struct {
	notificationRepo repository.Notifications
	centrifugoClient *centrifuge.Client
}

func NewNotificationUsecase(nr repository.Notifications, centrifugo *centrifuge.Client) *NotificationUsecase {
	return &NotificationUsecase{
		notificationRepo: nr,
		centrifugoClient: centrifugo,
	}
}

func (u *NotificationUsecase) GetAllNotifications(ctx context.Context, profileID uint) ([]models.Notification, error) {
	return u.notificationRepo.GetAllNotifications(ctx, profileID)
}

func (u *NotificationUsecase) MarkNotificationRead(ctx context.Context, id uint) error {
	return u.notificationRepo.MarkNotificationRead(ctx, id)
}

func (u *NotificationUsecase) SendOrderUpdateNotification(ctx context.Context, input models.UpdateOrderStatusPayload) error {
	payload := fmt.Sprintf(`{
		"type": "orderStatusChange",
		"data": {
			  "orderID": %d,
			  "oldStatus": "%s",
			  "newStatus": "%s"
		 }
	}`, input.OrderID, input.OldStatus, input.NewStatus)
	notification := models.CreateNotificationInput{
		Type:    "order_status_change",
		Payload: payload,
	}
	if err := u.notificationRepo.CreateNotification(ctx, input.OwnerID, notification); err != nil {
		return err
	}
	_, err := u.centrifugoClient.Publish(ctx, strconv.FormatUint(uint64(input.OwnerID), 10), []byte(payload))
	if err != nil {
		logger.Error(ctx, err.Error())
	}
	return err
}
