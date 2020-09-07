
package driver

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Phone string `bson:"phone"` 
    Email string `bson:"email"` 
  echo string
}

func New()*Driver{
  return &Driver{}
}
func (Driver *Driver)SetID(ID primitive.ObjectID)*Driver {
    Driver.echo = fmt.Sprintf("set ID from %v to %v",  Driver.ID, ID)
    Driver.ID=ID
    return Driver
}
func (Driver *Driver)SetName(Name string)*Driver {
    Driver.echo = fmt.Sprintf("set Name from %v to %v",  Driver.Name, Name)
    Driver.Name=Name
    return Driver
}
func (Driver *Driver)SetPhone(Phone string)*Driver {
    Driver.echo = fmt.Sprintf("set Phone from %v to %v",  Driver.Phone, Phone)
    Driver.Phone=Phone
    return Driver
}
func (Driver *Driver)SetEmail(Email string)*Driver {
    Driver.echo = fmt.Sprintf("set Email from %v to %v",  Driver.Email, Email)
    Driver.Email=Email
    return Driver
}
func (Driver *Driver)GetEcho()string {
    return Driver.echo
}