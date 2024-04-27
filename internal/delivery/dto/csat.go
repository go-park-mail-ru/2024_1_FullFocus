package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type PollStats struct {
	Title   string   `json:"title"`
	Rates   []uint32 `json:"rates"`
	Amount  uint32   `json:"amount"`
	Above70 uint32   `json:"above70"`
}

func ConvertStatsData(data models.PollStats) PollStats {
	return PollStats{
		data.Title,
		data.Rates,
		data.Amount,
		data.Above70,
	}
}

type Poll struct {
	ID    uint   `json:"ID"`
	Title string `json:"title"`
	Voted bool   `json:"voted"`
}

func ConvertPolls(polls []models.Poll) []Poll {
	dtoPolls := make([]Poll, 0, len(polls))
	for _, p := range polls {
		dtoPolls = append(dtoPolls, Poll{
			ID:    p.ID,
			Title: p.Title,
			Voted: p.Voted,
		})
	}
	return dtoPolls
}

type CreatePollRateInput struct {
	PollID uint `json:"pollID"`
	Rate   uint `json:"rate"`
}
