package hands

import (
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"net/http"
)

func Home(DB *mongo.Database, t *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("home")
		//w.WriteHeader(200)
		err := t.ExecuteTemplate(w, "index.html", struct{ Title string }{"HOME"})
		if OnError(w, err) {
			return
		}
	}
}
