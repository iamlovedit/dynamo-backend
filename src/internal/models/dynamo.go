package models

import "time"

type Dynamo struct {
	Id        string
	Name      string
	CreatedAt time.Time `json:"created_at"`
}
