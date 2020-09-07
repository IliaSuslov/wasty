
package driver

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
    
    
		
		Phone:=q.Get("Phone")
		if Phone!= ""{
		query[strings.ToLower("Phone")] = bson.M{"$regex": Phone}
		}
    
    
		
		Email:=q.Get("Email")
		if Email!= ""{
		query[strings.ToLower("Email")] = bson.M{"$regex": Email}
		}
    
		count, err := Col.CountDocuments(context.Background(), query)
		if OnError(w, err) {
			return
		}
		cur, err := Col.Find(context.Background(), query, GetOpts(q))
		if OnError(w, err) {
			return
		}
		vals := []*Driver{}
		w.Header().Set("ContentType", "application/json")
		for cur.Next(context.Background()) {
			val := &Driver{}
			err := cur.Decode(val)
			if OnError(w, err) {
				return
			}
			vals = append(vals, val)
		}
		err = json.NewEncoder(w).Encode(struct{Count int64; Items []*Driver}{count, vals})
		if OnError(w, err) {
			return
		}
	}
}
