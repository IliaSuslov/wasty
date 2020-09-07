
package driver

import (
	"github.com/alexsuslov/wasty/api/mid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)



func Route( r *mux.Router, prefix string,
	DB *mongo.Database,
	basicAuth mid.Middle,
	IsAllow func(User interface{}) bool,
	Denied mid.Response,
	None mid.Response,
	OnError func(http.ResponseWriter, error) bool){

	col:=DB.Collection("drivers")
	list := ListHDL(col, IsAllow, Denied, OnError)
	item := ItemHDL(col, IsAllow, Denied, None, OnError)
	update := UpdateHDL(col, IsAllow, Denied, None, OnError)

	r.HandleFunc(prefix+"/:id", basicAuth(item)).Methods("GET")
	r.HandleFunc(prefix+"/:id", basicAuth(update)).Methods("POST")
	r.HandleFunc(prefix+"s", basicAuth(list)).Methods("GET")
}