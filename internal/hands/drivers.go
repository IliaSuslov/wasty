package hands

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/wasty/api/model"
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type DriversQuery struct {
	Name   *string
	Limit  *int
	Offset *int
}

func Drivers(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := gcontext.Get(r, "user")
		User := user.(model.User)
		if !User.IsRole("admin", "manager"){
			Denied(w)
			return
		}
		query := bson.M{}
		drivers := DB.Collection("drivers")
		w.Header().Set("ContentType", "application/json")
		//todo: limit, skip
		cur, err := drivers.Find(context.Background(), query)
		if err != nil {
			log.Println(err)
		}
		res := []*model.Driver{}
		for cur.Next(context.Background()) {
			Driver := &model.Driver{}
			err := cur.Decode(Driver)
			if err != nil {
				log.Println(err)
			}
			res = append(res, Driver)
		}
		json.NewEncoder(w).Encode(res)
		if OnError(w, err) {
			return
		}
	}
}
