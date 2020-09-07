
package passwd

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Passwd struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Password [32]byte `bson:"password"` 
    Roles []string `bson:"roles"` 
    Firm interface{} `bson:"firm"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*Passwd{
  return &Passwd{}
}
func (Passwd *Passwd)SetID(ID primitive.ObjectID)*Passwd {
    Passwd.echo = fmt.Sprintf("set ID from %v to %v",  Passwd.ID, ID)
    Passwd.ID=ID
    return Passwd
}
func (Passwd *Passwd)SetName(Name string)*Passwd {
    Passwd.echo = fmt.Sprintf("set Name from %v to %v",  Passwd.Name, Name)
    Passwd.Name=Name
    return Passwd
}
func (Passwd *Passwd)SetPassword(Password [32]byte)*Passwd {
    Passwd.echo = fmt.Sprintf("set Password from %v to %v",  Passwd.Password, Password)
    Passwd.Password=Password
    return Passwd
}
func (Passwd *Passwd)SetRoles(Roles []string)*Passwd {
    Passwd.echo = fmt.Sprintf("set Roles from %v to %v",  Passwd.Roles, Roles)
    Passwd.Roles=Roles
    return Passwd
}
func (Passwd *Passwd)SetFirm(Firm interface{})*Passwd {
    Passwd.echo = fmt.Sprintf("set Firm from %v to %v",  Passwd.Firm, Firm)
    Passwd.Firm=Firm
    return Passwd
}
func (Passwd *Passwd)GetEcho()string {
    return Passwd.echo
}