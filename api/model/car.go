package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	Number string
	Trailer string
}