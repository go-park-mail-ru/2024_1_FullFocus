package models

type Poll struct {
	ID    uint
	Title string
	Voted bool
}

type CreatePollRate struct {
	ProfileID uint
	PollID    uint
	Rate      uint
}
