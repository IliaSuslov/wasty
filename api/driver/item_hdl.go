 
package driver
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

func ItemHDL(Col *mongo.Collection,
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
		ID, err := primitive.ObjectIDFromHex(id)
		if OnError(w, err) {
			return
		}
		query := bson.M{"_id": ID}
		result := &Driver{}
		err = Col.FindOne(context.Background(), query).Decode(result)
		if OnError(w, err) {
			return
		}
		w.Header().Set("ContentType", "application/json")
		json.NewEncoder(w).Encode(result)
		if OnError(w, err) {
			return
		}

	}
}
