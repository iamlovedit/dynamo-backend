package repository

import "github.com/iamlovedit/dynamo-backend/internal/models"

type DynamoRepository interface {
	FindById(id string) (*models.Dynamo, error)
}
