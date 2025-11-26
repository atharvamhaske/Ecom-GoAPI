package products

import (
	"context"
	"errors"

	repo "github.com/atharvamhaske/Ecom-GoAPI/internal/adapters/sqlc"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type CreateProductParams struct {
	Name         string `json:"name"`
	PriceInCents int64  `json:"priceInCents"`
	Quantity     int32  `json:"quantity"`
}

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	CreateProduct(ctx context.Context, params CreateProductParams) (repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) CreateProduct(ctx context.Context, params CreateProductParams) (repo.Product, error) {
	return s.repo.CreateProduct(ctx, repo.CreateProductParams{
		Name:           params.Name,
		PriceInCenters: int32(params.PriceInCents),
		Quantity:       params.Quantity,
	})
}

func (s *svc) GetProductByID(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.FindProductByID(ctx, id)
	if err != nil {
		return repo.Product{}, ErrProductNotFound
	}
	return product, nil
}
