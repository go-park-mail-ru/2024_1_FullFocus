package models

import "time"

type Notification struct {
	ID         uint
	Type       string
	ReadStatus string
	Payload    string
	CreatedAt  time.Time
}
