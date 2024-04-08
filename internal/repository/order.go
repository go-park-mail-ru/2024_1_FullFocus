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

func (r *OrderRepo) Create(ctx context.Context, userID uint, orderItems []models.OrderItem) (uint, error) {
	q := `INSERT INTO ordering (profile_id, order_status) VALUES (?, ?) RETURNING id;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", userID, "created")))
	start := time.Now()
	var orderID uint
	if err := r.storage.Get(ctx, &orderID, q, userID, "created"); err != nil {
		logger.Error(ctx, "error while creating order: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))

	q = `INSERT INTO order_item (ordering_id, product_id, count) VALUES (:ordering_id, :product_id, :count)`
	items := db.ConvertOrderItemsToTables(uint(orderID), orderItems)
	logger.Info(ctx, q, slog.Int("orders_amount", len(items)))
	start = time.Now()
	_, err := r.storage.NamedExec(ctx, q, items)
	if err != nil {
		logger.Error(ctx, "error while inserting order items: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))

	q = `UPDATE ordering
		 SET sum = (
		 	 SELECT SUM(p.price * o.count)
        	 FROM order_item o
                 INNER JOIN product p ON o.product_id = p.id
           	 WHERE o.ordering_id = 3
    	 )
		 WHERE id = 3;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start = time.Now()
	_, err = r.storage.Exec(ctx, q, orderID)
	if err != nil {
		logger.Error(ctx, "error while inserting sum: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))
	return uint(orderID), nil
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, orderID uint) (models.GetOrderPayload, error) {
	var orderProducts []db.OrderProduct
	q := `SELECT p.id, p.product_name, p.price, i.count, p.imgsrc
		  FROM order_item as i
			  INNER JOIN ordering AS o ON i.ordering_id = o.id
		      INNER JOIN product AS p ON i.product_id = p.id
		  WHERE o.id = ?`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start := time.Now()
	if err := r.storage.Get(ctx, &orderProducts, q, orderID); err != nil {
		logger.Error(ctx, "error while selecting order products: "+err.Error())
		return models.GetOrderPayload{}, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))

	var sum uint
	for _, product := range orderProducts {
		sum += product.Price
	}

	var orderInfo db.OrderInfo
	q = `SELECT order_status, DATE(created_at) FROM ordering WHERE id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start = time.Now()
	if err := r.storage.Get(ctx, &orderInfo, q, orderID); err != nil {
		logger.Error(ctx, "error while reading order status: "+err.Error())
		return models.GetOrderPayload{}, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return models.GetOrderPayload{
		Products:   db.ConvertOrderProductsToModels(orderProducts),
		Sum:        sum,
		Status:     orderInfo.Status,
		ItemsCount: uint(len(orderProducts)),
		CreatedAt:  orderInfo.CreatedAt,
	}, nil
}

func (r *OrderRepo) GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error) {
	q := `SELECT o.id, o.sum, o.order_status, count(i.product_id) AS items_count, DATE(o.created_at)
		  FROM ordering o
    	  	  LEFT JOIN order_item i ON o.id = i.ordering_id
		  WHERE o.profile_id = 3
		  GROUP BY o.id;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", profileID)))
	start := time.Now()
	var orders []db.Order
	if err := r.storage.Get(ctx, &orders, q, profileID); err != nil {
		logger.Error(ctx, "error while selecting orders: "+err.Error())
		return nil, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertOrdersFromTable(orders), nil
}

func (r *OrderRepo) GetProfileIDByOrderID(ctx context.Context, orderID uint) (uint, error) {
	q := `SELECT profile_id FROM ordering WHERE id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", orderID)))
	start := time.Now()
	var profileID uint
	if err := r.storage.Get(ctx, &profileID, q, orderID); err != nil {
		logger.Error(ctx, "error while selecting profile_id: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return profileID, nil
}

func (r *OrderRepo) Delete(ctx context.Context, orderID uint) error {
	q := `UPDATE ordering SET order_status = ? WHERE id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", orderID, "canceled")))
	start := time.Now()
	_, err := r.storage.Exec(ctx, q, orderID, "canceled")
	logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	if err != nil {
		logger.Error(ctx, "error while updating order status: "+err.Error())
		return err
	}
	return nil
}
