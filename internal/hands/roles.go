package hands

import (
	"context"
	"github.com/alexsuslov/wasty/api/model"
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"net/http"
)

func Roles(DB *mongo.Database, t *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		col := DB.Collection("roles")
		user := gcontext.Get(r, "user")
		User := user.(model.User)
		if !User.IsRole("admin") {
			Denied(w)
			return
		}
		q := r.URL.Query()
		query := bson.M{}
		queryAddStr := AddStr(query, q)
		queryAddRegex := AddLowerRegex(query, q)
		queryAddStr("ID")
		queryAddRegex("Name")

		Roles := []model.Role{}
		cur, err := col.Find(context.Background(), query, GetOpts(q))
		if OnError(w, err) {
			return
		}

		for cur.Next(context.Background()) {
			Role := model.Role{}
			err := cur.Decode(&Role)
			if OnError(w, err) {
				return
			}
			Roles = append(Roles, Role)
		}

		err = t.ExecuteTemplate(w, "roles.html", struct{ Title string }{"HOME"})
		if OnError(w, err) {
			return
		}
	}
}
