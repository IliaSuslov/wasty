
package waybill

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Waybill struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Number string `bson:"number"` 
    Дата string `bson:"lat"` 
    Desc string `bson:"desc"` 
    Car interface{} `bson:"car"` 
    Driver interface{} `bson:"driver"` 
    Firm interface{} `bson:"firm"` 
    Jobs []interface{} `bson:"jobs"` 
    Executions []interface{} `bson:"executions"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*Waybill{
  return &Waybill{}
}
func (Waybill *Waybill)SetID(ID primitive.ObjectID)*Waybill {
    Waybill.echo = fmt.Sprintf("set ID from %v to %v",  Waybill.ID, ID)
    Waybill.ID=ID
    return Waybill
}
func (Waybill *Waybill)SetNumber(Number string)*Waybill {
    Waybill.echo = fmt.Sprintf("set Number from %v to %v",  Waybill.Number, Number)
    Waybill.Number=Number
    return Waybill
}
func (Waybill *Waybill)SetДата(Дата string)*Waybill {
    Waybill.echo = fmt.Sprintf("set Дата from %v to %v",  Waybill.Дата, Дата)
    Waybill.Дата=Дата
    return Waybill
}
func (Waybill *Waybill)SetDesc(Desc string)*Waybill {
    Waybill.echo = fmt.Sprintf("set Desc from %v to %v",  Waybill.Desc, Desc)
    Waybill.Desc=Desc
    return Waybill
}
func (Waybill *Waybill)SetCar(Car interface{})*Waybill {
    Waybill.echo = fmt.Sprintf("set Car from %v to %v",  Waybill.Car, Car)
    Waybill.Car=Car
    return Waybill
}
func (Waybill *Waybill)SetDriver(Driver interface{})*Waybill {
    Waybill.echo = fmt.Sprintf("set Driver from %v to %v",  Waybill.Driver, Driver)
    Waybill.Driver=Driver
    return Waybill
}
func (Waybill *Waybill)SetFirm(Firm interface{})*Waybill {
    Waybill.echo = fmt.Sprintf("set Firm from %v to %v",  Waybill.Firm, Firm)
    Waybill.Firm=Firm
    return Waybill
}
func (Waybill *Waybill)SetJobs(Jobs []interface{})*Waybill {
    Waybill.echo = fmt.Sprintf("set Jobs from %v to %v",  Waybill.Jobs, Jobs)
    Waybill.Jobs=Jobs
    return Waybill
}
func (Waybill *Waybill)SetExecutions(Executions []interface{})*Waybill {
    Waybill.echo = fmt.Sprintf("set Executions from %v to %v",  Waybill.Executions, Executions)
    Waybill.Executions=Executions
    return Waybill
}
func (Waybill *Waybill)GetEcho()string {
    return Waybill.echo
}