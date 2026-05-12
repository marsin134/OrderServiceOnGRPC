package service

import (
	"context"
	"orderServiceGRPC/internal/models"
	"orderServiceGRPC/internal/repository"
)

type orderServiceInterface interface {
	CreateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	CetOrderId(ctx context.Context, id string) (*models.Order, error)
	GetOrderList(ctx context.Context, filter models.ListOrdersRequest) ([]*models.Order, error)
	UpdateOrder(ctx context.Context, data models.UpdateOrderData) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type Service struct {
	OrderService orderServiceInterface
}

func NewServiceGRPC(orderRepo repository.OrderRepository) *Service {
	orderServiceGRPC := NewOrderService(orderRepo)
	return &Service{
		OrderService: orderServiceGRPC,
	}
}
