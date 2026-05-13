package handler

import (
	"fmt"
	pb "orderServiceGRPC/pkg/generated/order"
	"strings"
)

func ValidateCreateOrderRequest(req *pb.CreateOrderRequest) error {
	var errors []string

	if strings.TrimSpace(req.ProductId) == "" {
		errors = append(errors, "product_id is required")
	}
	if strings.TrimSpace(req.SellerId) == "" {
		errors = append(errors, "seller_id is required")
	}
	if strings.TrimSpace(req.BuyerId) == "" {
		errors = append(errors, "buyer_id is required")
	}
	if strings.TrimSpace(req.PickupPoint) == "" {
		errors = append(errors, "pickup_point is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
