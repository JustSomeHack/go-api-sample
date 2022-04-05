package services

import (
	"context"

	"github.com/JustSomeHack/go-api-sample/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type DogsService interface {
	Add(ctx context.Context, dog *models.Dog) (*uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, filter interface{}) ([]models.Dog, error)
	GetOne(ctx context.Context, id uuid.UUID) (*models.Dog, error)
	Update(ctx context.Context, dog *models.Dog) error
}

type dogsService struct {
	db *gorm.DB
}

func NewDogsService(db *gorm.DB) DogsService {
	return &dogsService{
		db: db,
	}
}

func (s *dogsService) Add(ctx context.Context, dog *models.Dog) (*uuid.UUID, error) {
	if err := s.db.Create(dog).Error; err != nil {
		return nil, err
	}
	return &dog.ID, nil
}

func (s *dogsService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.db.Delete(models.Dog{}, "ID = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (s *dogsService) Get(ctx context.Context, filter interface{}) ([]models.Dog, error) {
	dogs := make([]models.Dog, 0)
	if err := s.db.Find(dogs).Error; err != nil {
		return nil, err
	}
	return dogs, nil
}

func (s *dogsService) GetOne(ctx context.Context, id uuid.UUID) (*models.Dog, error) {
	dog := new(models.Dog)
	if err := s.db.Find(dog, id).Error; err != nil {
		return nil, err
	}
	return dog, nil
}

func (s *dogsService) Update(ctx context.Context, dog *models.Dog) error {
	if err := s.db.Save(dog).Error; err != nil {
		return err
	}
	return nil
}
