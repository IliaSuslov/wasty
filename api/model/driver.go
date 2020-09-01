package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Driver struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	Phone string
	Email string
}
