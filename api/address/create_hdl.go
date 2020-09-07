
package address

import (
	"context"
	"encoding/json"
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func CreateHDL(Col *mongo.Collection,
	IsAllow func(User interface{}) bool,
	Denied func(w http.ResponseWriter),
	None func(w http.ResponseWriter),
	OnError func(http.ResponseWriter, error) bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := gcontext.Get(r, "user")
		if !IsAllow(user) {
			Denied(w)
			return
		}
		item := Address{}
		err := json.NewDecoder(r.Body).Decode(&item)
		if OnError(w, err) {
			return
		}
		res, err := Col.InsertOne(context.Background(), item)
		if OnError(w, err) {
			return
		}
		query:=bson.M{"_id":res.InsertedID}
		err = Col.FindOne(context.Background(), query).Decode(&item)
		if OnError(w, err) {
			return
		}
		w.Header().Set("ContentType", "application/json")
		json.NewEncoder(w).Encode(item)
		if OnError(w, err) {
			return
		}
	}
}
