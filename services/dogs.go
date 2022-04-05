package services

type DogsService interface {
}

type dogsService struct {
}

func NewDogsService() DogsService {
	return &dogsService{}
}
