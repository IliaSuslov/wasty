
package execution

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Execution struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Number string `bson:"name"` 
    Date time.Time `bson:"date"` 
    LoadingOn time.Time `bson:"loading_on"` 
    Category time.Time `bson:"category"` 
    Name time.Time `bson:"name"` 
    Address interface{} `bson:"address"` 
    Volume float32 `bson:"volume"` 
    Count float32 `bson:"count"` 
    Amount float32 `bson:"amount"` 
    Rest float32 `bson:"rest"` 
    Trash float32 `bson:"trash"` 
    Images []interface{} `bson:"images"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*Execution{
  return &Execution{}
}
func (Execution *Execution)SetID(ID primitive.ObjectID)*Execution {
    Execution.echo = fmt.Sprintf("set ID from %v to %v",  Execution.ID, ID)
    Execution.ID=ID
    return Execution
}
func (Execution *Execution)SetNumber(Number string)*Execution {
    Execution.echo = fmt.Sprintf("set Number from %v to %v",  Execution.Number, Number)
    Execution.Number=Number
    return Execution
}
func (Execution *Execution)SetDate(Date time.Time)*Execution {
    Execution.echo = fmt.Sprintf("set Date from %v to %v",  Execution.Date, Date)
    Execution.Date=Date
    return Execution
}
func (Execution *Execution)SetLoadingOn(LoadingOn time.Time)*Execution {
    Execution.echo = fmt.Sprintf("set LoadingOn from %v to %v",  Execution.LoadingOn, LoadingOn)
    Execution.LoadingOn=LoadingOn
    return Execution
}
func (Execution *Execution)SetCategory(Category time.Time)*Execution {
    Execution.echo = fmt.Sprintf("set Category from %v to %v",  Execution.Category, Category)
    Execution.Category=Category
    return Execution
}
func (Execution *Execution)SetName(Name time.Time)*Execution {
    Execution.echo = fmt.Sprintf("set Name from %v to %v",  Execution.Name, Name)
    Execution.Name=Name
    return Execution
}
func (Execution *Execution)SetAddress(Address interface{})*Execution {
    Execution.echo = fmt.Sprintf("set Address from %v to %v",  Execution.Address, Address)
    Execution.Address=Address
    return Execution
}
func (Execution *Execution)SetVolume(Volume float32)*Execution {
    Execution.echo = fmt.Sprintf("set Volume from %v to %v",  Execution.Volume, Volume)
    Execution.Volume=Volume
    return Execution
}
func (Execution *Execution)SetCount(Count float32)*Execution {
    Execution.echo = fmt.Sprintf("set Count from %v to %v",  Execution.Count, Count)
    Execution.Count=Count
    return Execution
}
func (Execution *Execution)SetAmount(Amount float32)*Execution {
    Execution.echo = fmt.Sprintf("set Amount from %v to %v",  Execution.Amount, Amount)
    Execution.Amount=Amount
    return Execution
}
func (Execution *Execution)SetRest(Rest float32)*Execution {
    Execution.echo = fmt.Sprintf("set Rest from %v to %v",  Execution.Rest, Rest)
    Execution.Rest=Rest
    return Execution
}
func (Execution *Execution)SetTrash(Trash float32)*Execution {
    Execution.echo = fmt.Sprintf("set Trash from %v to %v",  Execution.Trash, Trash)
    Execution.Trash=Trash
    return Execution
}
func (Execution *Execution)SetImages(Images []interface{})*Execution {
    Execution.echo = fmt.Sprintf("set Images from %v to %v",  Execution.Images, Images)
    Execution.Images=Images
    return Execution
}
func (Execution *Execution)GetEcho()string {
    return Execution.echo
}