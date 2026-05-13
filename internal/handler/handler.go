package handler

import (
	"orderServiceGRPC/internal/service"
	pb "orderServiceGRPC/pkg/generated/order"
)

type Handler struct {
	pb.UnimplementedOrderServiceServer

	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
