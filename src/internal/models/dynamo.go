package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dynamo struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	Created   time.Time          `json:"created"`
	Downloads int                `json:"downloads"`
	Votes     int                `json:"votes"`
	Updated   time.Time          `json:"latest_version_update"`
	Keywords  []string           `json:"keywords"`
	
}
