package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"

type Stat struct {
	Rate   uint `db:"rate"`
	Amount uint `db:"amount"`
}

func ConvertStatToModel(t Stat) models.StatRate {
	return models.StatRate{
		Rate:   t.Rate,
		Amount: t.Amount,
	}
}

func ConvertStatsToModels(tt []Stat) []models.StatRate {
	stats := make([]models.StatRate, 0)
	for _, t := range tt {
		stats = append(stats, ConvertStatToModel(t))
	}
	return stats
}
