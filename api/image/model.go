
package image

import (
  "fmt"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Image struct{
    ID primitive.ObjectID `bson:"_id,omitempty"` 
    Name string `bson:"name"` 
    Lat string `bson:"lat"` 
    Lng string `bson:"lng"` 
    Alt string `bson:"alt"` 
    ContentType string `bson:"content_type"` 
    Content []byte `bson:"content"` 
    echo string
    CreatedOn time.Ticker
    UpdatedOn time.Ticker
}

func New()*Image{
  return &Image{}
}
func (Image *Image)SetID(ID primitive.ObjectID)*Image {
    Image.echo = fmt.Sprintf("set ID from %v to %v",  Image.ID, ID)
    Image.ID=ID
    return Image
}
func (Image *Image)SetName(Name string)*Image {
    Image.echo = fmt.Sprintf("set Name from %v to %v",  Image.Name, Name)
    Image.Name=Name
    return Image
}
func (Image *Image)SetLat(Lat string)*Image {
    Image.echo = fmt.Sprintf("set Lat from %v to %v",  Image.Lat, Lat)
    Image.Lat=Lat
    return Image
}
func (Image *Image)SetLng(Lng string)*Image {
    Image.echo = fmt.Sprintf("set Lng from %v to %v",  Image.Lng, Lng)
    Image.Lng=Lng
    return Image
}
func (Image *Image)SetAlt(Alt string)*Image {
    Image.echo = fmt.Sprintf("set Alt from %v to %v",  Image.Alt, Alt)
    Image.Alt=Alt
    return Image
}
func (Image *Image)SetContentType(ContentType string)*Image {
    Image.echo = fmt.Sprintf("set ContentType from %v to %v",  Image.ContentType, ContentType)
    Image.ContentType=ContentType
    return Image
}
func (Image *Image)SetContent(Content []byte)*Image {
    Image.echo = fmt.Sprintf("set Content from %v to %v",  Image.Content, Content)
    Image.Content=Content
    return Image
}
func (Image *Image)GetEcho()string {
    return Image.echo
}