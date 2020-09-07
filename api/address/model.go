package address

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Address struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Lat       string             `bson:"lat"`
	Lng       string             `bson:"lng"`
	Alt       string             `bson:"alt"`
	echo      string
	CreatedOn time.Ticker
	UpdatedOn time.Ticker
}

func New() *Address {
	return &Address{}
}
func (Address *Address) SetID(ID primitive.ObjectID) *Address {
	Address.echo = fmt.Sprintf("set ID from %v to %v", Address.ID, ID)
	Address.ID = ID
	return Address
}
func (Address *Address) SetName(Name string) *Address {
	Address.echo = fmt.Sprintf("set Name from %v to %v", Address.Name, Name)
	Address.Name = Name
	return Address
}
func (Address *Address) SetLat(Lat string) *Address {
	Address.echo = fmt.Sprintf("set Lat from %v to %v", Address.Lat, Lat)
	Address.Lat = Lat
	return Address
}
func (Address *Address) SetLng(Lng string) *Address {
	Address.echo = fmt.Sprintf("set Lng from %v to %v", Address.Lng, Lng)
	Address.Lng = Lng
	return Address
}
func (Address *Address) SetAlt(Alt string) *Address {
	Address.echo = fmt.Sprintf("set Alt from %v to %v", Address.Alt, Alt)
	Address.Alt = Alt
	return Address
}
func (Address *Address) GetEcho() string {
	return Address.echo
}
