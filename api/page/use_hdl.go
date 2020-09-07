 
package page

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/wasty/api/passwd"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func PageHDL(Col *mongo.Collection,
	Denied func(w http.ResponseWriter),
	None func(w http.ResponseWriter),
	OnError func(http.ResponseWriter, error) bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name, ok := mux.Vars(r)["name"]
		if !ok {
			None(w)
			return
		}
		page := &Page{}
		query := bson.M{"name": name}
		err := Col.FindOne(context.Background(), query).Decode(page)
		if OnError(w, err) {
			return
		}
		var User *passwd.Passwd
		u := gcontext.Get(r, "user")
		if u!= nil{
			User = u.(*passwd.Passwd)
		}
		// if page.firm is set
		if page.Firm != nil && User.Firm !=nil && page.Firm != User.Firm {
			Denied(w)
			return
		}
		if len(page.Read)>0{
			if !User.IsRole(page.Read...){
				Denied(w)
				return
			}
		}


		w.Header().Set("ContentType", "application/json")
		json.NewEncoder(w).Encode(page)
		if OnError(w, err) {
			return
		}

	}
}
