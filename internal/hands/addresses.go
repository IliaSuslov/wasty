package hands

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/wasty/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type AddressQuery struct{
	Name *string
	Limit *int
	Offset *int
}

func Addresses(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//todo: check permishion
		query:=bson.M{}
		c := DB.Collection("addresses")
		w.Header().Set("ContentType", "application/json")
		//todo: limit, skip
		cur, err := c.Find(context.Background(), query)
		if err!= nil{
			log.Println(err)
		}
		for cur.Next(context.Background()){
			Addr := &model.Address{}
			err :=cur.Decode(Addr)
			if err!= nil{
				log.Println(err)
			}
			err=json.NewEncoder(w).Encode(Addr)
			if err!= nil{
				log.Println(err)
			}
		}
	}
}
