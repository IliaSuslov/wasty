
package car

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Car struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Number string `bson:"number"` 
    Trailer string `bson:"trailer"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*Car{
  return &Car{}
}
func (Car *Car)SetID(ID primitive.ObjectID)*Car {
    Car.echo = fmt.Sprintf("set ID from %v to %v",  Car.ID, ID)
    Car.ID=ID
    return Car
}
func (Car *Car)SetName(Name string)*Car {
    Car.echo = fmt.Sprintf("set Name from %v to %v",  Car.Name, Name)
    Car.Name=Name
    return Car
}
func (Car *Car)SetNumber(Number string)*Car {
    Car.echo = fmt.Sprintf("set Number from %v to %v",  Car.Number, Number)
    Car.Number=Number
    return Car
}
func (Car *Car)SetTrailer(Trailer string)*Car {
    Car.echo = fmt.Sprintf("set Trailer from %v to %v",  Car.Trailer, Trailer)
    Car.Trailer=Trailer
    return Car
}
func (Car *Car)GetEcho()string {
    return Car.echo
}