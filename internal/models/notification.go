package models

import "time"

type Notification struct {
	ID         uint
	Type       string
	ReadStatus bool
	Payload    string
	CreatedAt  time.Time
}
