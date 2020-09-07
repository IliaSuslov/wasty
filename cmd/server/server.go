package main

import (
	"context"
	"fmt"
	"github.com/alexsuslov/godotenv"
	"github.com/alexsuslov/wasty/api/address"
	"github.com/alexsuslov/wasty/api/car"
	"github.com/alexsuslov/wasty/api/driver"
	"github.com/alexsuslov/wasty/api/firm"
	"github.com/alexsuslov/wasty/api/page"
	"github.com/alexsuslov/wasty/api/passwd"
	"github.com/alexsuslov/wasty/api/user"
	"github.com/alexsuslov/wasty/api/waybill"
	"github.com/alexsuslov/wasty/internal/auth"
	"github.com/alexsuslov/wasty/internal/hands"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var version = "develop preview"

func main() {
	log.Println("Starting Monitor", version, "...")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Printf("warning .env not loaded")
	}
	//MONGO
	url := godotenv.GetPanic("MONGODB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("mongo connect error: %v", url)
		panic(err)
	}
	DB := client.Database(godotenv.GetPanic("DATABASE"))
	passwd.CreateAdmin(DB)

	//templates
	T, err := Load("templates", ".html")
	if err != nil {
		panic(err)
	}
	validator := auth.Validator(DB)
	basicAuth := auth.BasicAuth(validator)

	r := mux.NewRouter()
	r.HandleFunc("/", hands.Home(DB, T)).Methods("GET")
	IsAllow := func(User interface{})bool{
		return true
	}

	//drivers
	driver.Route(r, "/api/v1/driver", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)
	address.Route(r, "/api/v1/addr", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)
	car.Route(r, "/api/v1/car", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)
	firm.Route(r, "/api/v1/firm", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)
	user.Route(r, "/api/v1/user", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)
	waybill.Route(r, "/api/v1/waybill", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)
	page.Route(r, "/api/v1/page", DB, basicAuth, IsAllow, hands.Denied, hands.None, hands.OnError)

	view := page.PageHDL(DB.Collection("pages"), hands.Denied, hands.None, hands.OnError)
	r.HandleFunc("/api/v1/view/{name}", basicAuth(view)).Methods("GET")

	addr := fmt.Sprintf("%s:%d",
		godotenv.GetPanic("HTTP_HOST"),
		godotenv.GetIntPanic("HTTP_PORT"))

	staticDir := "/"
	FileServer := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/").
		Handler(http.StripPrefix(staticDir, FileServer))

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Web Listen", addr)
	log.Fatal(srv.ListenAndServe())
}

func Load(path string, ext string) (t *template.Template, err error) {
	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	var filesPath []string
	for _, fileInfo := range filesInfo {
		if !fileInfo.IsDir() {
			Ext := filepath.Ext(fileInfo.Name())
			if Ext == ext {
				filesPath = append(filesPath, path+"/"+fileInfo.Name())
			}
		}
	}
	return template.
		New(path).
		ParseFiles(filesPath...)
}
