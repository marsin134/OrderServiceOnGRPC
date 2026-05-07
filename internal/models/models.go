package models

import "time"

type Order struct {
	OrderId       string    `json:"orderId"`
	ProductId     string    `json:"productId"`
	SellerId      string    `json:"sellerId"`
	BuyerId       string    `json:"buyerId"`
	Status        string    `json:"status"`
	PickupPointId string    `json:"pickupPointId"`
	DeliveryTime  string    `json:"deliveryTime"`
	CreateTime    time.Time `json:"createTime"`
	UpdateTime    time.Time `json:"updateTime"`
}
