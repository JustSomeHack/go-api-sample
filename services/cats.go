package services

type CatsService interface {
}

type catsService struct {
}

func NewCatsService() CatsService {
	return &catsService{}
}
