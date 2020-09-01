package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	POI
}