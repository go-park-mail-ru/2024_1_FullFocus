package repository

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"

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

// CreatePromocodeItem и GetNewPromocode используются после создания заказа
// для определения доступного промокода и его выдачи.
func (r *PromocodeRepo) CreatePromocodeItem(ctx context.Context, info models.CreatePromocodeItemInput) error {
	q := `INSERT INTO promocode_item (promocode_type, profile_id, code) VALUES (?, ?, ?);`

	if _, err := r.storage.Exec(ctx, q, info.PromocodeID, info.ProfileID, info.Code); err != nil {
		logger.Error(ctx, err.Error())
		if pgErr := new(pgconn.PgError); errors.As(err, pgErr) {
			if strings.Contains(pgErr.Message, "duplicate") {
				return models.ErrPromocodeAlreadyExists
			}
		}
		return models.ErrInternal
	}
	return nil
}

func (r *PromocodeRepo) GetNewPromocode(ctx context.Context, sum uint) (uint, error) {
	q := `SELECT p.id
		  FROM promocode_item pi
				INNER JOIN promocode p ON pi.promocode_type = p.id
		  WHERE min_sum_give <= ?
		  ORDER BY min_sum_give DESC, min_sum_activation
		  LIMIT 1;`

	var newPromocodeID uint
	if err := r.storage.Get(ctx, &newPromocodeID, q, sum); err != nil {
		logger.Error(ctx, err.Error())
		return 0, models.ErrNoPromocode
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
	if err := r.storage.Select(ctx, &promocodeItem, q, pID, code); err != nil {
		logger.Error(ctx, err.Error())
		return models.PromocodeActivationTerms{}, models.ErrNoPromocode
	}
	if promocodeItem.ExpiresAt.Before(time.Now()) {
		return models.PromocodeActivationTerms{}, models.ErrPromocodeExpired
	}
	return dao.ConvertTerms(promocodeItem), nil
}

//func (r *PromocodeRepo) GetPromocodeOwnerID(ctx context.Context, promocodeID uint) (uint, error) {
//	q := `SELECT pi.profile_id
//		  FROM promocode_item pi
//		  WHERE pi.id = ?;`
//
//	var ownerID uint
//	if err := r.storage.Get(ctx, &ownerID, q, promocodeID); err != nil {
//		logger.Error(ctx, err.Error())
//		return 0, models.ErrNoPromocode
//	}
//	return ownerID, nil
//}

func (r *PromocodeRepo) GetAvailablePromocodes(ctx context.Context, profileID uint) ([]models.PromocodeItem, error) {
	q := `SELECT p.id,
			   p.name,
			   p.description,
			   pi.created_at + interval '1 hour' * p.ttl_hours - NOW() AS time_left,
			   pi.code,
			   p.min_sum_activation,
			   p.benefit_type,
			   p.value
		FROM promocode_item pi
			INNER JOIN promocode p ON pi.promocode_type = p.id
									AND profile_id = ?
		WHERE pi.created_at + interval '1 hour' * p.ttl_hours >= NOW();`

	var promocodes []dao.PromocodeItemInfo
	if err := r.storage.Select(ctx, &promocodes, q, profileID); err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrNoPromocode // handle properly
	}
	return dao.ConvertPromocodeItems(promocodes), nil
}

func (r *PromocodeRepo) ApplyPromocode(ctx context.Context, sum uint, promoID uint) (uint, error) {
	q := `SELECT p.benefit_type, p.value
		  FROM promocode_item pi
			INNER JOIN promocode p ON pi.promocode_type = p.id AND pi.id = ?;`

	var benefit dao.PromocodeBenefit
	if err := r.storage.Select(ctx, &benefit, q, promoID); err != nil {
		logger.Error(ctx, err.Error())
		if pgErr := new(pgconn.PgError); errors.As(err, pgErr) {
			if strings.Contains(pgErr.Message, "no rows") {
				return 0, models.ErrNoPromocode
			}
		}
		return 0, models.ErrInternal
	}
	switch benefit.Type {
	case "percentage":
		return sum * (100 - benefit.Value) / 100, nil
	case "price discount":
		return max(0, sum-benefit.Value), nil
	default:
		return sum, nil
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
