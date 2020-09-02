package hands

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/wasty/api/model"
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Atoin64(s string) *int64 {
	i, _ := strconv.Atoi(s)
	I := int64(i)
	return &I
}

func GetOpts(values url.Values) (opt *options.FindOptions) {
	sort := values.Get("sort")
	limit := values.Get("limit")
	skip := values.Get("skip")

	opt = &options.FindOptions{
		Limit: Atoin64(limit),
		Skip:  Atoin64(skip),
	}

	if strings.HasPrefix(sort, "!") {
		opt.Sort = bson.M{sort[1:]: -1}
	} else {
		opt.Sort = bson.M{sort: 1}
	}
	return
}

func Addresses(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		user := gcontext.Get(r, "user")
		User := user.(model.User)
		if !User.IsRole("admin", "manager") {
			Denied(w)
			return
		}

		//query
		q := r.URL.Query()
		query := bson.M{}
		qregex := AddLowerRegex(query, q)
		qregex("name")

		c := DB.Collection("addresses")
		w.Header().Set("ContentType", "application/json")

		ads := []*model.Address{}
		cur, err := c.Find(context.Background(), query, GetOpts(q))
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
