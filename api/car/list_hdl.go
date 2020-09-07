
package car

import (
	"context"
	"encoding/json"
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)


func Atoin64(s string, def int64) *int64 {
	if s ==""{
		return &def
	}
	i, _ := strconv.Atoi(s)
	I := int64(i)
	return &I
}


func GetOpts(values url.Values) (opt *options.FindOptions) {
	sort := values.Get("sort")
	limit := values.Get("limit")
	skip := values.Get("skip")

	opt = &options.FindOptions{
		Limit: Atoin64(limit, 100),
		Skip:  Atoin64(skip, 0),
	}
	if sort!= ""{
		if strings.HasPrefix(sort, "!") {
			opt.Sort = bson.M{sort[1:]: -1}
		} else {
			opt.Sort = bson.M{sort: 1}
		}
	}
	return
}


func ListHDL(Col *mongo.Collection,
	IsAllow func(User interface{}) bool, Denied func(w http.ResponseWriter),
	OnError func(http.ResponseWriter, error) bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := gcontext.Get(r, "user")
		if !IsAllow(user) {
			Denied(w)
			return
		}
		q := r.URL.Query()
		query := bson.M{}
    
    
		
		Name:=q.Get("Name")
		if Name!= ""{
		query[strings.ToLower("Name")] = bson.M{"$regex": Name}
		}
    
    
		
		Number:=q.Get("Number")
		if Number!= ""{
		query[strings.ToLower("Number")] = bson.M{"$regex": Number}
		}
    
    
		
		Trailer:=q.Get("Trailer")
		if Trailer!= ""{
		query[strings.ToLower("Trailer")] = bson.M{"$regex": Trailer}
		}
    

		cur, err := Col.Find(context.Background(), query, GetOpts(q))
		if OnError(w, err) {
			return
		}
		vals := []*Car{}
		w.Header().Set("ContentType", "application/json")
		for cur.Next(context.Background()) {
			val := &Car{}
			err := cur.Decode(val)
			if OnError(w, err) {
				return
			}
			vals = append(vals, val)
		}
		json.NewEncoder(w).Encode(vals)
		if OnError(w, err) {
			return
		}
	}
}
