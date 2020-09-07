
package page

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Page struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Data string `bson:"data"` 
    Type string `bson:"type"` 
    Firm *primitive.ObjectID `bson:"firm"` 
    Read []string `bson:"read"` 
    Write []string `bson:"write"` 
    echo string
    CreatedOn time.Time
    UpdatedOn time.Time
}

func New()*Page{
  return &Page{}
}
func (Page *Page)SetID(ID primitive.ObjectID)*Page {
    Page.echo = fmt.Sprintf("set ID from %v to %v",  Page.ID, ID)
    Page.ID=ID
    return Page
}
func (Page *Page)SetName(Name string)*Page {
    Page.echo = fmt.Sprintf("set Name from %v to %v",  Page.Name, Name)
    Page.Name=Name
    return Page
}
func (Page *Page)SetData(Data string)*Page {
    Page.echo = fmt.Sprintf("set Data from %v to %v",  Page.Data, Data)
    Page.Data=Data
    return Page
}
func (Page *Page)SetType(Type string)*Page {
    Page.echo = fmt.Sprintf("set Type from %v to %v",  Page.Type, Type)
    Page.Type=Type
    return Page
}
func (Page *Page)SetFirm(Firm *primitive.ObjectID)*Page {
    Page.echo = fmt.Sprintf("set Firm from %v to %v",  Page.Firm, Firm)
    Page.Firm=Firm
    return Page
}
func (Page *Page)SetRead(Read []string)*Page {
    Page.echo = fmt.Sprintf("set Read from %v to %v",  Page.Read, Read)
    Page.Read=Read
    return Page
}
func (Page *Page)SetWrite(Write []string)*Page {
    Page.echo = fmt.Sprintf("set Write from %v to %v",  Page.Write, Write)
    Page.Write=Write
    return Page
}
func (Page *Page)GetEcho()string {
    return Page.echo
}