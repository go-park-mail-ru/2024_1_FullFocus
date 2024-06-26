package models

type Poll struct {
	ID    uint
	Title string
	Voted bool
}

type CreatePollRateInput struct {
	ProfileID uint
	PollID    uint
	Rate      uint
}

type PollStats struct {
	Title   string
	Rates   []uint32
	Amount  uint32
	Above70 uint32
}
