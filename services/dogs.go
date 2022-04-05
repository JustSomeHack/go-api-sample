package services

import (
	"context"

	"github.com/JustSomeHack/go-api-sample/models"

	"github.com/google/uuid"
)

type DogsService interface {
	Add(ctx context.Context, cat *models.Dog) (*uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, filter interface{}) ([]models.Cat, error)
	GetOne(ctx context.Context, id uuid.UUID) (*models.Cat, error)
	Update(ctx context.Context, id uuid.UUID, cat *models.Cat) error
}

type dogsService struct {
}

func NewDogsService() DogsService {
	return &dogsService{}
}

func (s *dogsService) Add(ctx context.Context, cat *models.Dog) (*uuid.UUID, error) {
	return nil, nil
}

func (s *dogsService) Delete(ctx context.Context, id uuid.UUID) error {

	return nil
}

func (s *dogsService) Get(ctx context.Context, filter interface{}) ([]models.Cat, error) {
	return nil, nil
}

func (s *dogsService) GetOne(ctx context.Context, id uuid.UUID) (*models.Cat, error) {
	return nil, nil
}

func (s *dogsService) Update(ctx context.Context, id uuid.UUID, cat *models.Cat) error {
	return nil
}
