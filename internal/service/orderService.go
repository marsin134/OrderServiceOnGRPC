package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"orderServiceGRPC/internal/models"
	"orderServiceGRPC/internal/repository"
	pb "orderServiceGRPC/pkg/generated/order"
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

func (svc *orderService) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order := models.Order{
		OrderId:       uuid.New().String(),
		ProductId:     request.ProductId,
		SellerId:      request.SellerId,
		BuyerId:       request.BuyerId,
		Status:        OrderStatusName[1],
		PickupPointId: request.PickupPoint,
		DeliveryTime:  request.EstimatedDeliveryTime,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	_, err := svc.order.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: modelToProto(&order)}, nil
}

func (svc *orderService) CetOrderId(ctx context.Context, id string) (*pb.OrderResponse, error) {
	order, err := svc.order.GetOrderId(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: modelToProto(order)}, err
}

func (svc *orderService) GetOrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	filter := models.ListOrdersRequest{
		IdProduct:    req.IdProduct,
		IdSeller:     req.IdSeller,
		IdBuyer:      req.IdBuyer,
		IdPickup:     req.IdPickupPoint,
		DeliveryTime: req.DeliveryTime,
	}

	if req.Status != nil {
		status, err := protoStatusToModel(*req.Status)
		if err != nil {
			return nil, err
		}
		filter.Status = &status
	}

	orders, err := svc.order.GetOrderList(ctx, filter)
	if err != nil {
		return nil, err
	}

	ordersListPb := make([]*pb.Order, len(orders))
	for i, order := range orders {
		ordersListPb[i] = modelToProto(order)
	}
	return &pb.OrderListResponse{Orders: ordersListPb}, nil
}

func (svc *orderService) UpdateOrder(ctx context.Context, data *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error) {
	order, err := svc.order.GetOrderId(ctx, data.OrderId)
	if err != nil {
		return nil, err
	}

	if data.NewStatus == nil {
		orderNextIndex := OrderStatusIndex[order.Status]
		orderNextIndex++

		order.Status = OrderStatusName[orderNextIndex]
	} else {
		newStatusString, err := protoStatusToModel(*data.NewStatus)
		if err != nil {
			return nil, err
		}
		order.Status = newStatusString
	}

	order.UpdateTime = time.Now()
	newOrder, err := svc.order.UpdateOrder(ctx, *order)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{Order: modelToProto(newOrder)}, nil
}

func (svc *orderService) DeleteOrder(ctx context.Context, orderId string) error {
	return svc.order.DeleteOrder(ctx, orderId)
}

func modelToProto(order *models.Order) *pb.Order {
	return &pb.Order{
		OrderId:               order.OrderId,
		ProductId:             order.ProductId,
		SellerId:              order.SellerId,
		BuyerId:               order.BuyerId,
		Status:                order.Status,
		PickupPointId:         order.PickupPointId,
		EstimatedDeliveryTime: order.DeliveryTime,
		CreatedAt:             order.CreateTime.Format(time.RFC3339),
		UpdatedAt:             order.UpdateTime.Format(time.RFC3339),
	}
}

func protoStatusToModel(protoStatus pb.OrderStatus) (string, error) {
	switch protoStatus {
	case pb.OrderStatus_ORDER_STATUS_UNSPECIFIED:
		return OrderStatusName[0], nil

	case pb.OrderStatus_ORDER_STATUS_COLLECTING:
		return OrderStatusName[1], nil

	case pb.OrderStatus_ORDER_STATUS_DELIVERING:
		return OrderStatusName[2], nil

	case pb.OrderStatus_ORDER_STATUS_READY_FOR_PICKUP:
		return OrderStatusName[3], nil

	case pb.OrderStatus_ORDER_STATUS_DELIVERED:
		return OrderStatusName[4], nil

	case pb.OrderStatus_ORDER_STATUS_RETURNED:
		return OrderStatusName[5], nil

	default:
		return "", fmt.Errorf("unknown order status: %v", protoStatus)
	}
}
