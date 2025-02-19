package repository

import (
	"context"

	"github.com/iamlovedit/dynamo-backend/internal/config"
	"github.com/iamlovedit/dynamo-backend/internal/models"
	"github.com/iamlovedit/dynamo-backend/internal/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgodb "go.mongodb.org/mongo-driver/mongo"
)

type MongoDynamoRepository struct {
	collection *mgodb.Collection
}

func NewMongoDynamoRepository() *MongoDynamoRepository {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	client, err := mongo.ConnectDB(cfg.MongoURI)
	if err != nil {
		panic(err)
	}

	collection := mongo.GetCollection(client, "dynamos")
	return &MongoDynamoRepository{
		collection: collection,
	}
}

func (r *MongoDynamoRepository) FindById(id string) (*models.Dynamo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var dynamo models.Dynamo
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&dynamo)
	if err != nil {
		return nil, err
	}

	return &dynamo, nil
}
