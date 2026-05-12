package service

import (
	"context"
	"github.com/google/uuid"
	"orderServiceGRPC/internal/models"
	"orderServiceGRPC/internal/repository"
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

var (
	OrderStatusIndex = map[string]int{
		"ORDER_STATUS_UNSPECIFIED":      0,
		"ORDER_STATUS_COLLECTING":       1,
		"ORDER_STATUS_DELIVERING":       2,
		"ORDER_STATUS_READY_FOR_PICKUP": 3,
		"ORDER_STATUS_DELIVERED":        4,
		"ORDER_STATUS_RETURNED":         5,
	}
)

type orderService struct {
	order repository.OrderRepository
}

func NewOrderService(order repository.OrderRepository) *orderService {
	return &orderService{
		order: order,
	}
}

func (svc *orderService) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	order.OrderId = uuid.New().String()
	order.Status = OrderStatusName[0]
	order.CreateTime = time.Now()
	order.UpdateTime = time.Now()

	return svc.order.CreateOrder(ctx, order)
}

func (svc *orderService) CetOrderId(ctx context.Context, id string) (*models.Order, error) {
	return svc.order.GetOrderId(ctx, id)
}

func (svc *orderService) GetOrderList(ctx context.Context, filter models.ListOrdersRequest) ([]*models.Order, error) {
	return svc.order.GetOrderList(ctx, filter)
}

func (svc *orderService) UpdateOrder(ctx context.Context, data models.UpdateOrderData) (*models.Order, error) {
	order, err := svc.order.GetOrderId(ctx, data.OrderId)
	if err != nil {
		return nil, err
	}

	if data.OrderStatusIndex == 0 {
		orderNextIndex := OrderStatusIndex[order.Status]
		orderNextIndex++

		order.Status = OrderStatusName[orderNextIndex]
	} else {
		order.Status = OrderStatusName[data.OrderStatusIndex]
	}

	order.UpdateTime = time.Now()

	return svc.order.UpdateOrder(ctx, *order)
}

func (svc *orderService) DeleteOrder(ctx context.Context, orderId string) error {
	return svc.order.DeleteOrder(ctx, orderId)
}
