package service

import (
	"context"
	"orderServiceGRPC/internal/repository"
	pb "orderServiceGRPC/pkg/generated/order"
)

type orderServiceInterface interface {
	CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.OrderResponse, error)
	CetOrderId(ctx context.Context, id string) (*pb.OrderResponse, error)
	GetOrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error)
	UpdateOrder(ctx context.Context, data *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error)
	DeleteOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderDeleteResponse, error)
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
