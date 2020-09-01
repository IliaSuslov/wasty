package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Firm struct {
	ID primitive.ObjectID
	Name string
	Fullname string
	Phone string
	Email string
	OwnerID string
}