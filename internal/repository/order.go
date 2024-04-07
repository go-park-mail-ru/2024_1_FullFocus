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
	createOrderingQuery    = `INSERT INTO ozon.ordering (profile_id, order_status) VALUES (?, ?) RETURNING id;`
	insertOrderItemsQuery  = `INSERT INTO ozon.order_item (ordering_id, product_id, count) VALUES (:ordering_id, :product_id, :count);`
	updateOrderStatusQuery = `UPDATE ozon.ordering SET order_status = ? WHERE id = ?;`
	selectProfileID        = `SELECT profile_id FROM ozon.ordering WHERE id = ?;`
	selectOrderProducts    = `SELECT p.product_name, p.price, i.count, p.imgsrc FROM ozon.order_item as i INNER JOIN ozon.ordering AS o ON i.ordering_id = o.id INNER JOIN ozon.product AS p ON i.product_id = p.id WHERE o.id = ?`
	selectOrderStatus      = `SELECT order_status FROM ozon.ordering WHERE id = ?;`
	selectAllOrdersInfo    = `SELECT id, sum, order_status FROM ozon.ordering WHERE profile_id = ?;`
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
	orderID, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, "error while creating order: "+err.Error())
		return 0, err
	}
	items := db.ConvertOrderItemsToTables(uint(orderID), orderItems)
	logger.Info(ctx, createOrderingQuery, slog.Int("orders_amount", len(items)))
	start = time.Now()
	_, err = r.storage.NamedExec(ctx, insertOrderItemsQuery, items)
	if err != nil {
		logger.Error(ctx, "error while inserting order items: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))
	return uint(orderID), nil
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, orderID uint) (models.GetOrderPayload, error) {
	var orderProducts []db.OrderProduct
	logger.Info(ctx, selectOrderProducts, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start := time.Now()
	if err := r.storage.Get(ctx, &orderProducts, selectOrderProducts, orderID); err != nil {
		logger.Error(ctx, "error while selecting order products: "+err.Error())
		return models.GetOrderPayload{}, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))

	var sum uint
	for _, product := range orderProducts {
		sum += product.Price
	}

	var status string
	logger.Info(ctx, selectOrderStatus, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start = time.Now()
	if err := r.storage.Get(ctx, &status, selectOrderStatus, orderID); err != nil {
		logger.Error(ctx, "error while reading order status: "+err.Error())
		return models.GetOrderPayload{}, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return models.GetOrderPayload{
		Products: db.ConvertOrderProductsToModels(orderProducts),
		Sum:      sum,
		Status:   status,
	}, nil
}

func (r *OrderRepo) GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error) {
	logger.Info(ctx, selectAllOrdersInfo, slog.String("args", fmt.Sprintf("$1 = %d", profileID)))
	start := time.Now()
	var orders []db.Order
	if err := r.storage.Get(ctx, &orders, selectAllOrdersInfo, profileID); err != nil {
		logger.Error(ctx, "error while selecting orders: "+err.Error())
		return nil, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertOrdersFromTable(orders), nil
}

func (r *OrderRepo) GetProfileIDByOrderingID(ctx context.Context, orderID uint) (uint, error) {
	logger.Info(ctx, selectProfileID, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start := time.Now()
	var profileID uint
	if err := r.storage.Get(ctx, &profileID, createOrderingQuery, orderID); err != nil {
		logger.Error(ctx, "error while selecting profile_id: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return profileID, nil
}

func (r *OrderRepo) Delete(ctx context.Context, orderID uint) error {
	logger.Info(ctx, updateOrderStatusQuery, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", orderID, "canceled")))
	start := time.Now()
	_, err := r.storage.Exec(ctx, updateOrderStatusQuery, orderID, "canceled")
	logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	if err != nil {
		logger.Error(ctx, "error while updating order status: "+err.Error())
		return err
	}
	return nil
}
