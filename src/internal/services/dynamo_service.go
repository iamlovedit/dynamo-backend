package services

import (
	"github.com/iamlovedit/dynamo-backend/internal/models"
	"github.com/iamlovedit/dynamo-backend/internal/repository"
)

type DynamoService struct {
	repo repository.DynamoRepository
}

func (s *DynamoService) FindDynamo(id string) (*models.Dynamo, error) {
	return s.repo.FindById(id)
}
