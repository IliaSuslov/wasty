
package firm

import (
	"github.com/alexsuslov/wasty/api/mid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
type handler func(w http.ResponseWriter, r *http.Request)

func Route( r *mux.Router, prefix string,
	DB *mongo.Database,
	basicAuth mid.Middle,
	IsAllow func(User interface{}) bool,
	Denied func(w http.ResponseWriter),
	None func(w http.ResponseWriter),
	OnError func(http.ResponseWriter, error) bool){

	col:=DB.Collection("firms")
	list := ListHDL(col, IsAllow, Denied, OnError)
	item := ItemHDL(col, IsAllow, Denied, None, OnError)
	update := UpdateHDL(col, IsAllow, Denied, None, OnError)
	create := CreateHDL(col, IsAllow, Denied, None, OnError)

	r.HandleFunc(prefix+"/:id", basicAuth(item)).Methods("GET")
	r.HandleFunc(prefix+"/:id", basicAuth(update)).Methods("POST")
	r.HandleFunc(prefix+"s", basicAuth(list)).Methods("GET")
	r.HandleFunc(prefix+"s", basicAuth(create)).Methods("POST")
}
