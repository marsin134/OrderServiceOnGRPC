package repository

import (
	"context"
	"orderServiceGRPC/internal/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	GetOrderId(ctx context.Context, id string) (*models.Order, error)
	GetOrderList(ctx context.Context, filter models.ListOrdersRequest) ([]*models.Order, error)
	UpdateOrder(ctx context.Context, id string, newStatusIndex int) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type Repository struct {
	Order OrderRepository
}

func NewRepository(order OrderRepository) *Repository {
	return &Repository{
		Order: order,
	}
}
