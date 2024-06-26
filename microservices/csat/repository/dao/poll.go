package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"

type PollTable struct {
	ID    uint   `db:"id"`
	Title string `db:"title"`
	Voted int    `db:"voted"`
}

func ConvertPollTableToModel(t PollTable) models.Poll {
	var votedFlag bool
	if t.Voted == 1 {
		votedFlag = true
	}
	return models.Poll{
		ID:    t.ID,
		Title: t.Title,
		Voted: votedFlag,
	}
}

func ConvertPollTablesToModels(tt []PollTable) []models.Poll {
	polls := make([]models.Poll, 0)
	for _, t := range tt {
		polls = append(polls, ConvertPollTableToModel(t))
	}
	return polls
}
