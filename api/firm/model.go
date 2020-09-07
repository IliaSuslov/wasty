
package firm

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Firm struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Contact string `bson:"contact"` 
    Users []interface{} `bson:"users"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*Firm{
  return &Firm{}
}
func (Firm *Firm)SetID(ID primitive.ObjectID)*Firm {
    Firm.echo = fmt.Sprintf("set ID from %v to %v",  Firm.ID, ID)
    Firm.ID=ID
    return Firm
}
func (Firm *Firm)SetName(Name string)*Firm {
    Firm.echo = fmt.Sprintf("set Name from %v to %v",  Firm.Name, Name)
    Firm.Name=Name
    return Firm
}
func (Firm *Firm)SetContact(Contact string)*Firm {
    Firm.echo = fmt.Sprintf("set Contact from %v to %v",  Firm.Contact, Contact)
    Firm.Contact=Contact
    return Firm
}
func (Firm *Firm)SetUsers(Users []interface{})*Firm {
    Firm.echo = fmt.Sprintf("set Users from %v to %v",  Firm.Users, Users)
    Firm.Users=Users
    return Firm
}
func (Firm *Firm)GetEcho()string {
    return Firm.echo
}