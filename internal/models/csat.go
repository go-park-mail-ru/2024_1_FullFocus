package models

type Poll struct {
	ID    uint
	Title string
	Voted bool
}

type GetPollsInput struct {
	ProfileID uint
}

type CreatePollRateInput struct {
	ProfileID uint
	PollID    uint
	Rate      uint
}

type PollStats struct {
	Stats        StatsData
	PrimaryStats StatsData
}

type StatsData struct {
	Rates   []uint32
	Amount  uint32
	Above70 uint32
}
