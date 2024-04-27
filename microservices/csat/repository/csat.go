package repository

import (
	"context"
	"errors"
	"strings"

	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/repository/dao"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
	"github.com/jackc/pgx/v5/pgconn"
)

type CSATRepo struct {
	storage db.Database
}

func NewCSATRepo(st db.Database) *CSATRepo {
	return &CSATRepo{
		storage: st,
	}
}

func (r *CSATRepo) GetAllPolls(ctx context.Context, profileID uint) ([]models.Poll, error) {
	q := `SELECT p.id, p.title,
	CASE
		WHEN r.profile_id IS NULL THEN 0
		ELSE 1
	END AS voted
	FROM poll p
	LEFT JOIN response r ON p.id = r.poll_id AND r.profile_id = $1;`

	polls := make([]dao.PollTable, 0)
	if err := r.storage.Select(ctx, &polls, q, profileID); err != nil {
		logger.Error(ctx, err.Error())
		return []models.Poll{}, commonError.ErrInternal
	}
	if len(polls) == 0 {
		return []models.Poll{}, commonError.ErrNotFound
	}
	return dao.ConvertPollTablesToModels(polls), nil
}

func (r *CSATRepo) CreatePollRate(ctx context.Context, input models.CreatePollRate) error {
	q := `INSERT INTO response(poll_id, profile_id, rate) VALUES($1, $2, $3);`

	if _, err := r.storage.Exec(ctx, q, input.PollID, input.ProfileID, input.Rate); err != nil {
		logger.Error(ctx, err.Error())
		var sqlErr *pgconn.PgError
		if errors.As(err, &sqlErr) {
			if strings.Contains(sqlErr.Message, "duplicate") {
				return commonError.ErrAlreadyExists
			}
			if strings.Contains(sqlErr.Message, "foreign") {
				return commonError.ErrNotFound
			}
		}
		return commonError.ErrInternal
	}
	return nil
}

func (r *CSATRepo) GetPollStats(ctx context.Context, pollID uint) (string, []models.StatRate, error) {
	q := `SELECT p.title
	FROM poll p
	WHERE p.id = $1;`

	var title string
	if err := r.storage.Get(ctx, &title, q, pollID); err != nil {
		logger.Error(ctx, err.Error())
		return "", []models.StatRate{}, commonError.ErrInternal
	}
	if title == "" {
		return "", []models.StatRate{}, commonError.ErrNotFound
	}

	q = `SELECT r.rate, count(*) AS amount
	FROM response r
	WHERE r.poll_id = $1
	GROUP BY r.rate;`

	stats := make([]dao.Stat, 0)
	if err := r.storage.Select(ctx, &stats, q, pollID); err != nil {
		logger.Error(ctx, err.Error())
		return "", []models.StatRate{}, commonError.ErrInternal
	}
	if len(stats) == 0 {
		return "", []models.StatRate{}, commonError.ErrNotFound
	}

	return title, dao.ConvertStatsToModels(stats), nil
}
