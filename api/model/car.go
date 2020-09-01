package model


import (
	//gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID primitive.ObjectID
	Name string
	Number string
	Trailer string
}