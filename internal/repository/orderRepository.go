package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"orderServiceGRPC/internal/database"
	"orderServiceGRPC/internal/models"
	"time"
)

var (
	OrderStatusName = map[int]string{
		0: "ORDER_STATUS_UNSPECIFIED",
		1: "ORDER_STATUS_COLLECTING",
		2: "ORDER_STATUS_DELIVERING",
		3: "ORDER_STATUS_READY_FOR_PICKUP",
		4: "ORDER_STATUS_DELIVERED",
		5: "ORDER_STATUS_RETURNED",
	}
)

type orderRepository struct {
	db *database.DB
}

func NewOrderRepository(db *database.DB) *orderRepository {
	return &orderRepository{db}
}

func (r orderRepository) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	order.OrderId = uuid.New().String()
	order.Status = OrderStatusName[0]
	order.CreateTime = time.Now()
	order.UpdateTime = time.Now()

	query := `
		INSERT INTO orders (order_id, product_id, seller_id, buyer_id, status, pickup_point_id, delivery_time, created_at, updated_at)
		VALUES (:order_id, :product_id, :seller_id, :buyer_id, :status, :pickup_point_id, :delivery_time, :created_at, :updated_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, order)
	if err != nil {
		return nil, fmt.Errorf("error when creating a order when accessing the database: %w", err)
	}

	return &order, nil
}

func (r orderRepository) GetOrderId(ctx context.Context, id string) (*models.Order, error) {
	query := `SELECT * FROM orders WHERE order_id = $1`

	var order models.Order
	err := r.db.GetContext(ctx, &order, query, id)
	if err != nil {
		return nil, fmt.Errorf("error when getting the order id from the database: %w", err)
	}

	return &order, nil
}

func (r orderRepository) GetOrderList(ctx context.Context, filter models.ListOrdersRequest) ([]*models.Order, error) {
	baseQuery := `
        FROM orders
        WHERE 1=1
    `

	var args []interface{}
	argCount := 0

	// Dynamic addition of filtering conditions
	if filter.IdBuyer != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND buyer_id = $%d", argCount)
		args = append(args, *filter.IdBuyer)
	}

	if filter.IdSeller != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND seller_id = $%d", argCount)
		args = append(args, *filter.IdSeller)
	}

	if filter.IdProduct != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND product_id = $%d", argCount)
		args = append(args, *filter.IdProduct)
	}

	if filter.IdPickup != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND pickup_point_id = $%d", argCount)
		args = append(args, *filter.IdPickup)
	}

	if filter.Status != nil {
		argCount++
		baseQuery += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *filter.Status)
	}

	query := "SELECT * " + baseQuery

	var orders []*models.Order
	err := r.db.SelectContext(ctx, &orders, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error when getting the orders when accessing the database: %w", err)
	}
	return orders, nil
}

func (r orderRepository) UpdateOrder(ctx context.Context, id string, newStatusIndex int) (*models.Order, error) {
	order, err := r.GetOrderId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error when getting the order id from the database: %w", err)
	}
	order.Status = OrderStatusName[newStatusIndex]

	query := `UPDATE orders SET status = :status 
              WHERE order_id = :order_id`

	_, err = r.db.NamedExecContext(ctx, query, order)
	if err != nil {
		return nil, fmt.Errorf("error when updating the order status from the database: %w", err)
	}

	return order, nil
}

func (r orderRepository) DeleteOrder(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE order_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error when deleting the order from the database: %w", err)
	}
	return nil
}
