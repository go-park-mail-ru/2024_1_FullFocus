package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

type OrderRepo struct {
	storage db.Database
}

func NewOrderRepo(dbClient db.Database) *OrderRepo {
	return &OrderRepo{
		storage: dbClient,
	}
}

const (
	createOrderingQuery    = `INSERT INTO ordering (profile_id, order_status) VALUES (?, ?) RETURNING ordering_id;`
	insertOrderItemsQuery  = `INSERT INTO order_item (ordering_id, product_id, count) VALUES (:ordering_id, :product_id, :count);`
	updateOrderStatusQuery = `UPDATE ordering SET order_status = ? WHERE id = ?;`
	selectProfileID        = `SELECT profile_id FROM ordering WHERE id = ?;`
	selectOrderProducts    = `SELECT p.products_name, i.count, p.imgsrc FROM order_item as i INNER JOIN ordering AS o ON i.ordering_id = o.id INNER JOIN product AS p ON i.product_id = p.id WHERE o.id = ?;`
)

func (r *OrderRepo) Create(ctx context.Context, userID uint, orderItems []models.OrderItem) (uint, error) {
	logger.Info(ctx, createOrderingQuery, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", userID, "created")))
	start := time.Now()
	result, err := r.storage.Exec(ctx, createOrderingQuery, userID, "created")
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))
	if err != nil {
		logger.Error(ctx, "error while creating order: "+err.Error())
		return 0, err
	}
	orderingID, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, "error while creating order: "+err.Error())
		return 0, err
	}
	items := db.ConvertOrderItemsToTables(uint(orderingID), orderItems)
	logger.Info(ctx, createOrderingQuery, slog.Int("orders_amount", len(items)))
	start = time.Now()
	_, err = r.storage.NamedExec(ctx, insertOrderItemsQuery, items)
	if err != nil {
		logger.Error(ctx, "error while inserting order items: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))
	return uint(orderingID), nil
}

func (r *OrderRepo) GetOrderProducts(ctx context.Context, orderingID uint) ([]models.OrderProduct, error) {
	logger.Info(ctx, selectOrderProducts, slog.String("args", fmt.Sprintf("$1 = %d", orderingID)))
	var orderProducts []db.OrderProduct
	start := time.Now()
	if err := r.storage.Get(ctx, &orderProducts, selectOrderProducts, orderingID); err != nil {
		logger.Error(ctx, "error while selecting order products: "+err.Error())
		return nil, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertOrderProductsToModels(orderProducts), nil
}

func (r *OrderRepo) GetProfileIDByOrderingID(ctx context.Context, orderingID uint) (uint, error) {
	logger.Info(ctx, selectProfileID, slog.String("args", fmt.Sprintf("$1 = %d", orderingID)))
	start := time.Now()
	var profileID uint
	if err := r.storage.Get(ctx, &profileID, createOrderingQuery, orderingID); err != nil {
		logger.Error(ctx, "error while selecting profile_id: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return profileID, nil
}

func (r *OrderRepo) Delete(ctx context.Context, orderingID uint) error {
	logger.Info(ctx, updateOrderStatusQuery, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", orderingID, "canceled")))
	start := time.Now()
	_, err := r.storage.Exec(ctx, updateOrderStatusQuery, orderingID, "canceled")
	logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	if err != nil {
		logger.Error(ctx, "error while updating order status: "+err.Error())
		return err
	}
	return nil
}
