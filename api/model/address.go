package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressesQuery struct {
	Name  *string
	Skip  *int
	Limit *int
}

type IAddress interface {
	Addresses(AddressesQuery) ([]*Address, error)
	AddressByID(ID string) (*Address, error)
	UpdateAddress(*Address) (*Address, error)
	DeleteAddress(ID string) error
}

type Address struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	POI
}

type CAddress struct {
	DB *mongo.Database
}

func NewCAddress(db *mongo.Database) *CAddress {
	return &CAddress{
		db,
	}
}
