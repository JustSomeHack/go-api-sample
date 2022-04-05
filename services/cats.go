package services

import (
	"context"

	"github.com/JustSomeHack/go-api-sample/models"

	"github.com/google/uuid"
)

type CatsService interface {
	Add(ctx context.Context, cat *models.Cat) (*uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, filter interface{}) ([]models.Cat, error)
	GetOne(ctx context.Context, id uuid.UUID) (*models.Cat, error)
	Update(ctx context.Context, id uuid.UUID, cat *models.Cat) error
}

type catsService struct {
}

func NewCatsService() CatsService {
	return &catsService{}
}

func (s *catsService) Add(ctx context.Context, cat *models.Cat) (*uuid.UUID, error) {
	return nil, nil
}

func (s *catsService) Delete(ctx context.Context, id uuid.UUID) error {

	return nil
}

func (s *catsService) Get(ctx context.Context, filter interface{}) ([]models.Cat, error) {
	return nil, nil
}

func (s *catsService) GetOne(ctx context.Context, id uuid.UUID) (*models.Cat, error) {
	return nil, nil
}

func (s *catsService) Update(ctx context.Context, id uuid.UUID, cat *models.Cat) error {
	return nil
}
