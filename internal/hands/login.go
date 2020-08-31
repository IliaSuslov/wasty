package hands

import (
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"net/http"
)

func Login(DB *mongo.Database, t *template.Template)func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){

	}
}