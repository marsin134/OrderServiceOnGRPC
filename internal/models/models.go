package models

import "time"

type Order struct {
	OrderId       string    `json:"orderId" db:"order_id"`
	ProductId     string    `json:"productId" db:"product_id"`
	SellerId      string    `json:"sellerId" db:"seller_id"`
	BuyerId       string    `json:"buyerId" db:"buyer_id"`
	Status        string    `json:"status" db:"status"`
	PickupPointId string    `json:"pickupPointId" db:"pickup_point_id"`
	DeliveryTime  string    `json:"deliveryTime" db:"delivery_time"`
	CreateTime    time.Time `json:"createTime" db:"created_at"`
	UpdateTime    time.Time `json:"updateTime" db:"updated_at"`
}

type ListOrdersRequest struct {
	IdProduct *string
	IdSeller  *string
	IdBuyer   *string
	IdPickup  *string
	Status    *string
}

type UpdateOrderData struct {
	OrderId          string
	OrderStatusIndex int //optional
}
