package services

import (
	"context"

	"github.com/JustSomeHack/go-api-sample/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CatsService interface {
	Add(ctx context.Context, cat *models.Cat) (*uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, filter interface{}) ([]models.Cat, error)
	GetOne(ctx context.Context, id uuid.UUID) (*models.Cat, error)
	Update(ctx context.Context, cat *models.Cat) error
}

type catsService struct {
	db *gorm.DB
}

func NewCatsService(db *gorm.DB) CatsService {
	return &catsService{
		db: db,
	}
}

func (s *catsService) Add(ctx context.Context, cat *models.Cat) (*uuid.UUID, error) {
	if err := s.db.Create(cat).Error; err != nil {
		return nil, err
	}
	return &cat.ID, nil
}

func (s *catsService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.db.Delete(models.Cat{}, "ID = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (s *catsService) Get(ctx context.Context, filter interface{}) ([]models.Cat, error) {
	cats := make([]models.Cat, 0)
	if err := s.db.Find(cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (s *catsService) GetOne(ctx context.Context, id uuid.UUID) (*models.Cat, error) {
	cat := new(models.Cat)
	if err := s.db.Find(cat, id).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *catsService) Update(ctx context.Context, cat *models.Cat) error {
	if err := s.db.Save(cat).Error; err != nil {
		return err
	}
	return nil
}
