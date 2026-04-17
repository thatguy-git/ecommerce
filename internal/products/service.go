package products

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/thatguy-git/ecom/internal/adapters/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductById(ctx context.Context, id pgtype.UUID) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func Newservice(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) FindProductById(ctx context.Context, id pgtype.UUID) (repo.Product, error) {
	return s.repo.FindProductById(ctx, id)
}
