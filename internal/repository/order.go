package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type OrderRepo struct {
	storage db.Database
}

func NewOrderRepo(dbClient db.Database) *OrderRepo {
	return &OrderRepo{
		storage: dbClient,
	}
}

func (r *OrderRepo) Create(ctx context.Context, userID uint, sum uint, orderItems []models.OrderItem) (uint, error) {
	tx, err := r.storage.Begin(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = errors.Join(err, rollbackErr)
			}
		}
	}()
	q := `INSERT INTO ordering (profile_id, order_status) VALUES ($1, $2) RETURNING id;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", userID, "created")))
	start := time.Now()
	var orderID uint
	if err = tx.GetContext(ctx, &orderID, q, userID, "created"); err != nil {
		logger.Error(ctx, "error while creating order: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))

	q = `INSERT INTO order_item (ordering_id, product_id, count) VALUES (:ordering_id, :product_id, :count)`
	items := dao.ConvertOrderItemsToTables(orderID, orderItems)
	logger.Info(ctx, q, slog.Int("orders_amount", len(items)))
	start = time.Now()
	_, err = tx.NamedExecContext(ctx, q, items)
	if err != nil {
		logger.Error(ctx, "error while inserting order items: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))

	q = `UPDATE ordering
		 SET sum = $1
		 WHERE id = $2;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start = time.Now()
	_, err = tx.ExecContext(ctx, q, sum, orderID)
	if err != nil {
		logger.Error(ctx, "error while inserting sum: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return orderID, err
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, orderID uint) (models.GetOrderPayload, error) {
	var orderProducts []dao.OrderProduct
	q := `SELECT p.id, p.product_name, p.price, i.count, p.imgsrc
		  FROM order_item as i
			  INNER JOIN ordering AS o ON i.ordering_id = o.id
		      INNER JOIN product AS p ON i.product_id = p.id
		  WHERE o.id = ?`
	if err := r.storage.Select(ctx, &orderProducts, q, orderID); err != nil {
		logger.Error(ctx, "error while selecting order products: "+err.Error())
		return models.GetOrderPayload{}, models.ErrNoRowsFound
	}

	var sum uint
	for _, product := range orderProducts {
		sum += product.Price * product.Count
	}

	var orderInfo dao.OrderInfo
	q = `SELECT order_status, DATE(created_at) AS created_at FROM ordering WHERE id = ?;`
	if err := r.storage.Get(ctx, &orderInfo, q, orderID); err != nil {
		logger.Error(ctx, "error while reading order status: "+err.Error())
		return models.GetOrderPayload{}, models.ErrNoRowsFound
	}
	return models.GetOrderPayload{
		Products:   dao.ConvertOrderProductsToModels(orderProducts),
		Sum:        sum,
		Status:     orderInfo.Status,
		ItemsCount: uint(len(orderProducts)),
		CreatedAt:  orderInfo.CreatedAt,
	}, nil
}

func (r *OrderRepo) GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error) {
	q := `SELECT o.id, o.sum, o.order_status, count(i.product_id) AS items_count, DATE(o.created_at) AS created_at
		  FROM ordering o
    	  	  LEFT JOIN order_item i ON o.id = i.ordering_id
		  WHERE o.profile_id = ?
		  GROUP BY o.id
		  ORDER BY o.id DESC;`
	var orders []dao.Order
	if err := r.storage.Select(ctx, &orders, q, profileID); err != nil {
		logger.Error(ctx, "error while selecting orders: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertOrdersFromTable(orders), nil
}

func (r *OrderRepo) GetProfileIDByOrderID(ctx context.Context, orderID uint) (uint, error) {
	q := `SELECT profile_id FROM ordering WHERE id = ?;`
	var profileID uint
	if err := r.storage.Get(ctx, &profileID, q, orderID); err != nil {
		logger.Error(ctx, "error while selecting profile_id: "+err.Error())
		return 0, models.ErrNoRowsFound
	}
	return profileID, nil
}

func (r *OrderRepo) Delete(ctx context.Context, orderID uint) error {
	q := `UPDATE ordering SET order_status = 'cancelled' WHERE id = ?;`
	_, err := r.storage.Exec(ctx, q, orderID)
	if err != nil {
		logger.Error(ctx, "error while updating order status: "+err.Error())
		return err
	}
	return nil
}
