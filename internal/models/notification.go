package models

import "time"

type Notification struct {
	ID         uint
	Type       string
	ReadStatus bool
	Payload    string
	CreatedAt  time.Time
}

type OrderNotification struct {
	ProfileID uint
	OrderID   uint
	OldStatus string
	NewStatus string
}
