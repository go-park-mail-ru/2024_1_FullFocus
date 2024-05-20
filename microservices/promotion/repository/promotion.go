package repository

import (
	"context"
	"errors"
	"strings"

	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/repository/dao"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repo struct {
	storage db.Database
}

func NewRepo(st db.Database) *Repo {
	return &Repo{
		storage: st,
	}
}

func (r *Repo) CreatePromoProductInfo(ctx context.Context, input models.PromoData) error {
	q := `INSERT INTO promo_product(product_id, benefit_type, benefit_value) VALUES(?, ?, ?);`

	if _, err := r.storage.Exec(ctx, q, input.ProductID, input.BenefitType, input.BenefitValue); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		var sqlErr *pgconn.PgError
		if errors.As(err, &sqlErr) {
			if strings.Contains(sqlErr.Message, "duplicate") {
				return models.ErrPromoProductAlreadyExists
			}
			if strings.Contains(sqlErr.Message, "foreign") {
				return models.ErrProductNotFound
			}
		}
		return commonError.ErrInternal
	}
	return nil
}

func (r *Repo) GetAllPromoProductsIDs(ctx context.Context) ([]uint, error) {
	q := `SELECT product_id FROM promo_product;`

	prIDs := make([]uint, 0)
	if err := r.storage.Select(ctx, &prIDs, q); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return nil, commonError.ErrInternal
	}
	if len(prIDs) == 0 {
		return nil, models.ErrProductNotFound
	}
	return prIDs, nil
}

func (r *Repo) GetPromoProductsInfoByIDs(ctx context.Context, prIDs []uint) ([]models.PromoData, error) {
	q := `SELECT product_id, benefit_type, benefit_value
	FROM promo_product
	WHERE product_id = ANY(?)
	ORDER BY array_position(?, product_id);`

	promoData := make([]dao.PromoProductTable, 0, len(prIDs))
	if err := r.storage.Select(ctx, &promoData, q, prIDs, prIDs); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return nil, commonError.ErrInternal
	}
	if len(promoData) == 0 {
		return nil, models.ErrProductNotFound
	}
	return dao.ConvertPromoProductTablesToModels(promoData), nil
}

func (r *Repo) DeletePromoProductInfo(ctx context.Context, prID uint) error {
	q := `DELETE FROM promo_product WHERE product_id = ? RETURNING product_id;`

	var id uint
	if err := r.storage.Get(ctx, &id, q, prID); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return models.ErrProductNotFound
	}
	return nil
}
