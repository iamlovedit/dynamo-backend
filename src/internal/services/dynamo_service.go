package services

import (
	"github.com/iamlovedit/dynamo-backend/internal/models"
	"github.com/iamlovedit/dynamo-backend/internal/repository"
)

type IDynamoService interface {
	FindDynamo(id string) (*models.Dynamo, error)
}

type DynamoService struct {
	repo repository.DynamoRepository
}

func NewDynamoService(repo repository.DynamoRepository) IDynamoService {
	return &DynamoService{
		repo: repo,
	}
}

func (s *DynamoService) FindDynamo(id string) (*models.Dynamo, error) {
	return s.repo.FindById(id)
}
