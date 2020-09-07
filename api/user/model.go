
package user

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type User struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Fullname string `bson:"fname"` 
    Firm interface{} `bson:"firm"` 
    Email string `bson:"email"` 
    QRcode string `bson:"qrcode"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*User{
  return &User{}
}
func (User *User)SetID(ID primitive.ObjectID)*User {
    User.echo = fmt.Sprintf("set ID from %v to %v",  User.ID, ID)
    User.ID=ID
    return User
}
func (User *User)SetName(Name string)*User {
    User.echo = fmt.Sprintf("set Name from %v to %v",  User.Name, Name)
    User.Name=Name
    return User
}
func (User *User)SetFullname(Fullname string)*User {
    User.echo = fmt.Sprintf("set Fullname from %v to %v",  User.Fullname, Fullname)
    User.Fullname=Fullname
    return User
}
func (User *User)SetFirm(Firm interface{})*User {
    User.echo = fmt.Sprintf("set Firm from %v to %v",  User.Firm, Firm)
    User.Firm=Firm
    return User
}
func (User *User)SetEmail(Email string)*User {
    User.echo = fmt.Sprintf("set Email from %v to %v",  User.Email, Email)
    User.Email=Email
    return User
}
func (User *User)SetQRcode(QRcode string)*User {
    User.echo = fmt.Sprintf("set QRcode from %v to %v",  User.QRcode, QRcode)
    User.QRcode=QRcode
    return User
}
func (User *User)GetEcho()string {
    return User.echo
}