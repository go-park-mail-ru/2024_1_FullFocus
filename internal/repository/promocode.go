package repository

import (
	"context"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type PromocodeRepo struct {
	storage db.Database
}

func NewPromocodeRepo(dbClient db.Database) *PromocodeRepo {
	return &PromocodeRepo{
		storage: dbClient,
	}
}

func (r *PromocodeRepo) CreatePromocodeItem(ctx context.Context, info models.CreatePromocodeItemInput) error {
	q := `INSERT INTO promocode_item (promocode_type, profile_id, code) VALUES (?, ?, ?);`

	if _, err := r.storage.Exec(ctx, q, info.PromocodeID, info.ProfileID, info.Code); err != nil {
		logger.Error(ctx, err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return models.ErrPromocodeAlreadyExists
		}
		return models.ErrInternal
	}
	return nil
}

func (r *PromocodeRepo) GetNewPromocode(ctx context.Context, sum uint) (uint, error) {
	q := `SELECT p.id
		  FROM promocode p
		  WHERE min_sum_give <= ?
		  ORDER BY min_sum_give DESC, RANDOM()
		  LIMIT 1;`

	var newPromocodeID uint
	if err := r.storage.Get(ctx, &newPromocodeID, q, sum); err != nil {
		logger.Error(ctx, err.Error())
		if strings.Contains(err.Error(), "no rows") {
			return 0, models.ErrNoPromocode
		}
		return 0, models.ErrInternal
	}
	return newPromocodeID, nil
}

func (r *PromocodeRepo) GetPromocodeItemByActivationCode(ctx context.Context, pID uint, code string) (models.PromocodeActivationTerms, error) {
	q := `SELECT pi.id,
				pi.created_at + interval '1 hour' * p.ttl_hours AS expires_at,
				p.min_sum_activation,
				p.benefit_type,
				p.value
			FROM promocode_item pi
				INNER JOIN promocode p ON pi.promocode_type = p.id 
										AND pi.profile_id = ?
										AND pi.code = ?;`

	var promocodeItem dao.PromocodeActivationTerms
	if err := r.storage.Get(ctx, &promocodeItem, q, pID, code); err != nil {
		logger.Error(ctx, err.Error())
		if strings.Contains(err.Error(), "no rows") {
			return models.PromocodeActivationTerms{}, models.ErrNoPromocode
		}
		return models.PromocodeActivationTerms{}, models.ErrInternal
	}
	if promocodeItem.ExpiresAt.Before(time.Now()) {
		return models.PromocodeActivationTerms{}, models.ErrPromocodeExpired
	}
	return dao.ConvertTerms(promocodeItem), nil
}

func (r *PromocodeRepo) GetPromocodeByID(ctx context.Context, promocodeID uint) (models.Promocode, error) {
	q := `SELECT p.id, p.description, p.min_sum_give, p.min_sum_activation, p.benefit_type, p.value, p.ttl_hours
		  FROM promocode p
		  WHERE p.id = ?;`

	var promocode dao.Promocode
	if err := r.storage.Get(ctx, &promocode, q, promocodeID); err != nil {
		logger.Error(ctx, err.Error())
		if strings.Contains(err.Error(), "no rows") {
			return models.Promocode{}, models.ErrNoPromocode
		}
		return models.Promocode{}, models.ErrInternal
	}
	return dao.ConvertPromocode(promocode), nil
}

func (r *PromocodeRepo) GetAvailablePromocodes(ctx context.Context, profileID uint) ([]models.PromocodeItem, error) {
	q := `SELECT pi.id,
			   p.description,
			   pi.created_at + interval '1 hour' * p.ttl_hours - NOW() AS time_left,
			   pi.code,
			   p.min_sum_activation,
			   p.benefit_type,
			   p.value
		FROM promocode_item pi
			INNER JOIN promocode p ON pi.promocode_type = p.id
									AND profile_id = ?
		WHERE pi.created_at + interval '1 hour' * p.ttl_hours >= NOW()
		ORDER BY pi.created_at DESC;`

	var promocodes []dao.PromocodeItemInfo
	if err := r.storage.Select(ctx, &promocodes, q, profileID); err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrInternal
	}
	if len(promocodes) == 0 {
		return nil, models.ErrNoPromocode
	}
	return dao.ConvertPromocodeItems(promocodes), nil
}

func (r *PromocodeRepo) GetPromocodesAmount(ctx context.Context, profileID uint) (uint, error) {
	q := `SELECT count(*)
		  FROM promocode_item pi
			INNER JOIN promocode p ON pi.promocode_type = p.id
				AND profile_id = ?
		  WHERE pi.created_at + interval '1 hour' * p.ttl_hours >= NOW();`

	var amount uint
	if err := r.storage.Get(ctx, &amount, q, profileID); err != nil {
		logger.Error(ctx, err.Error())
		return 0, models.ErrNoPromocode
	}
	return amount, nil
}

func (r *PromocodeRepo) ApplyPromocode(ctx context.Context, input models.ApplyPromocodeInput) (uint, error) {
	q := `SELECT p.min_sum_activation,
       			p.benefit_type,
       			p.value,
       			pi.created_at + interval '1 hour' * p.ttl_hours AS expires_at
		  FROM promocode_item pi
			INNER JOIN promocode p ON pi.promocode_type = p.id AND pi.id = ? AND pi.profile_id = ?;`

	var benefit dao.PromocodeBenefit
	if err := r.storage.Get(ctx, &benefit, q, input.PromoID, input.ProfileID); err != nil {
		logger.Error(ctx, err.Error())
		if strings.Contains(err.Error(), "no rows") {
			return 0, models.ErrNoPromocode
		}
		return 0, models.ErrInternal
	}
	if input.Sum < benefit.MinSum {
		return 0, models.ErrInvalidPromocode
	}
	if benefit.ExpiresAt.Before(time.Now()) {
		return 0, models.ErrPromocodeExpired
	}
	switch benefit.Type {
	case "percentage":
		return input.Sum * (100 - benefit.Value) / 100, nil
	case "price discount":
		return max(0, input.Sum-benefit.Value), nil
	default:
		return input.Sum, nil
	}
}

func (r *PromocodeRepo) DeleteUsedPromocode(ctx context.Context, id uint) error {
	q := `DELETE FROM promocode_item pi
		  WHERE pi.id = ?;`

	if _, err := r.storage.Exec(ctx, q, id); err != nil {
		logger.Error(ctx, err.Error())
		return models.ErrInternal
	}
	return nil
}
