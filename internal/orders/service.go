package orders

import (
	"context"
	"errors"
	"fmt"
	"time"

	repo "github.com/atharvamhaske/Ecom-GoAPI/internal/adapters/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
	ErrOrderNotFound   = errors.New("order not found")
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	// validate payload
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer ID is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least one item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	// create an order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	// look for the product if exists
	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNoStock
		}

		// create order item
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: int64(product.PriceInCenters),
		})
		if err != nil {
			return repo.Order{}, err
		}

		// Challenge: Update the product stock quantity
		newQuantity := product.Quantity - item.Quantity
		_, err = qtx.UpdateProductStock(ctx, repo.UpdateProductStockParams{
			ID:       item.ProductID,
			Quantity: newQuantity,
		})
		if err != nil {
			return repo.Order{}, err
		}
	}

	tx.Commit(ctx)
	return order, nil
}

func (s *svc) GetOrderByID(ctx context.Context, orderID int64) (OrderResponse, error) {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return OrderResponse{}, ErrOrderNotFound
	}

	orderItems, err := s.repo.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return OrderResponse{}, err
	}

	var items []OrderItemResponse
	var totalPriceCents int64 = 0

	for _, item := range orderItems {
		product, err := s.repo.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return OrderResponse{}, ErrProductNotFound
		}

		itemTotal := int64(item.Quantity) * item.PriceCents
		totalPriceCents += itemTotal

		items = append(items, OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			PriceCents:  item.PriceCents,
			TotalCents:  itemTotal,
		})
	}

	var createdAt time.Time
	if order.CreatedAt.Valid {
		createdAt = order.CreatedAt.Time
	}

	return OrderResponse{
		ID:              order.ID,
		CustomerID:      order.CustomerID,
		TotalPriceCents: totalPriceCents,
		Items:           items,
		CreatedAt:       createdAt,
	}, nil
}
