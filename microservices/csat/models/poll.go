package models

type Poll struct {
	ID    uint
	Title string
}

type CreatePollRate struct {
	ProfileID uint
	PollID    uint
	Rate      uint
}
