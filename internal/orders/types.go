package orders

import (
	"context"
	"time"

	repo "github.com/atharvamhaske/Ecom-GoAPI/internal/adapters/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type OrderItemResponse struct {
	ProductID   int64 `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int32 `json:"quantity"`
	PriceCents  int64 `json:"priceCents"`
	TotalCents  int64 `json:"totalCents"`
}

type OrderResponse struct {
	ID              int64              `json:"id"`
	CustomerID      int64              `json:"customerId"`
	TotalPriceCents int64              `json:"totalPriceCents"`
	Items           []OrderItemResponse `json:"items"`
	CreatedAt       time.Time          `json:"createdAt"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
	GetOrderByID(ctx context.Context, orderID int64) (OrderResponse, error)
}

