package hands

import (
	"context"
	"github.com/alexsuslov/wasty/api/model"
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Roles(DB *mongo.Database, t *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		col := DB.Collection("roles")
		roles := gcontext.Get(r, "user_roles")
		if model.IsRole(roles, "admin") == -1 {
			w.WriteHeader(401)
			return
		}
		q := r.URL.Query()
		query := bson.M{}
		queryAddStr := AddStr(query, q)
		queryAddRegex := AdderLowerRegex(query, q)

		sort := "!_id"
		if q.Get("sort") != "" {
			sort = q.Get("sort")
		}

		queryAddStr("ID")
		queryAddRegex("Name")

		Limit := int64(100)
		opt := options.FindOptions{
			Limit: &Limit,
		}
		if strings.HasPrefix(sort, "!") {
			opt.Sort = bson.M{sort[1:]: -1}
		} else {
			opt.Sort = bson.M{sort: 1}
		}
		limit, err := strconv.Atoi(q.Get("limit"))
		if err==nil{
			l:=int64(limit)
			opt.Limit=&l
		}
		Roles:=[]model.Role{}
		cur, err := col.Find(context.Background(), query, &opt)
		if err != nil{
			//todo: 500 error page
			log.Println("Find:", err)
			return
		}

		for cur.Next(context.Background()){
			Role:=model.Role{}
			err:=cur.Decode(&Role)
			if err!=nil{
				//todo: 500 error page
				log.Printf("Decode:%v, raw:%#v",err, cur.Current.String())
				continue
			}
			Roles = append(Roles, Role)
		}

		err = t.ExecuteTemplate(w, "roles.html", struct{ Title string }{"HOME"})
		if err != nil {
			//todo: 500 error page
			log.Println(err)
		}
	}
}
