package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"orderServiceGRPC/internal/config"
	pb "orderServiceGRPC/pkg/generated/order"
)

func (h *Handler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	config.Log.Info("Create order called",
		zap.String("id product", req.ProductId))

	if err := ValidateCreateOrderRequest(req); err != nil {
		config.Log.Warn("invalid create order request", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	order, err := h.service.OrderService.CreateOrder(ctx, req)
	if err != nil {
		config.Log.Error("Failed to create order", zap.Error(err))
		return nil, handleServiceError(err)
	}

	config.Log.Info("Create order successfully", zap.Any("order", order))

	return order, nil
}

func (h *Handler) GetOrderId(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	config.Log.Info("Get order called", zap.String("id", req.Id))

	order, err := h.service.OrderService.CetOrderId(ctx, req.Id)
	if err != nil {
		config.Log.Error("Failed to get order", zap.Error(err))
		return nil, handleServiceError(err)
	}

	config.Log.Info("Get order successfully")
	return order, nil
}

func (h *Handler) GetOrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	config.Log.Info("Get order list called")

	orderList, err := h.service.OrderService.GetOrderList(ctx, req)
	if err != nil {
		config.Log.Error("Failed to get order list", zap.Error(err))
		return nil, handleServiceError(err)
	}
	config.Log.Info("Get order list successfully")
	return orderList, nil
}

func (h *Handler) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error) {
	config.Log.Info("Get update order called")

	order, err := h.service.OrderService.UpdateOrder(ctx, req)
	if err != nil {
		config.Log.Error("Failed to update order", zap.Error(err))
		return nil, handleServiceError(err)
	}
	config.Log.Info("Update order successfully")
	return order, nil
}

func (h *Handler) DeleteOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderDeleteResponse, error) {
	config.Log.Info("Delete order called")

	response, err := h.service.OrderService.DeleteOrder(ctx, req)
	if err != nil {
		config.Log.Error("Failed to delete order", zap.Error(err))
		return nil, handleServiceError(err)
	}

	config.Log.Info("Delete order successfully")
	return response, nil
}
