
package page

import (
	"context"
	"encoding/json"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Patch struct{
	OP string
	Path string
	Values []interface{}
}

func UpdateHDL(Col *mongo.Collection,
	IsAllow func(User interface{}) bool, Denied func(w http.ResponseWriter), None func(w http.ResponseWriter),
	OnError func(http.ResponseWriter, error) bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		user := gcontext.Get(r, "user")

		if !IsAllow(user) {
			Denied(w)
			return
		}
		id, ok := mux.Vars(r)["id"]
		if !ok {
			None(w)
			return
		}
		item := Page{}
		err := json.NewDecoder(r.Body).Decode(&item)
		if OnError(w, err) {
			return
		}
		ID, err := primitive.ObjectIDFromHex(id)
		if OnError(w, err) {
			return
		}
		query := bson.M{"_id": ID}
		_, err = Col.UpdateOne(context.Background(), query, bson.M{"$set":item})
		if OnError(w, err) {
			return
		}
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
