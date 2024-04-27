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

func (r *CSATRepo) GetAllPolls(ctx context.Context) ([]models.Poll, error) {
	q := `SELECT id, title FROM poll;`

	polls := make([]dao.PollTable, 0)
	if err := r.storage.Select(ctx, &polls, q); err != nil {
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
