package hands

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/wasty/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strings"
)

func GetOpts(sort string, limit *int64, skip *int64) (opt *options.FindOptions) {
	opt = &options.FindOptions{
		Limit: limit,
		Skip:  skip,
	}

	if strings.HasPrefix(sort, "!") {
		opt.Sort = bson.M{sort[1:]: -1}
	} else {
		opt.Sort = bson.M{sort: 1}
	}
	return
}

func OnError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	w.WriteHeader(500)
	err = json.NewEncoder(w).Encode(struct {
		Success bool
		Message string
	}{
		false,
		err.Error(),
	})
	if err != nil {
		log.Println(err)
	}
	return true
}

func Addresses(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//todo: check permishion

		//query
		q := &struct {
			Name  *string
			Skip  *int64
			Limit *int64
		}{}
		json.NewDecoder(r.Body).Decode(&q)
		defer r.Body.Close()
		//name
		query := bson.M{}
		if q.Name != nil {
			query["name"] = bson.M{"$regex": q.Name}
		}

		c := DB.Collection("addresses")
		w.Header().Set("ContentType", "application/json")
		opt := &options.FindOptions{
			Limit: q.Limit,
			Skip:  q.Skip,
		}
		ads := []*model.Address{}
		cur, err := c.Find(context.Background(), query, opt)
		if OnError(w, err) {
			return
		}
		for cur.Next(context.Background()) {
			Addr := &model.Address{}
			err := cur.Decode(Addr)
			if OnError(w, err) {
				return
			}
			ads = append(ads, Addr)
		}
		json.NewEncoder(w).Encode(ads)
		if OnError(w, err) {
			return
		}
	}
}
