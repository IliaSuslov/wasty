package hands

import (
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"net/http"
)

func Users(DB *mongo.Database, t *template.Template)func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		roles:=r.Header.Get("w_roles")
		log.Println("users", roles)
		//w.WriteHeader(200)
		err:=t.ExecuteTemplate(w, "users.html", struct{Title string}{"HOME"})
		if err!= nil{
			log.Println(err)
		}
	}
}