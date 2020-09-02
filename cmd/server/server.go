package main

import (
	"context"
	"fmt"
	"github.com/alexsuslov/godotenv"

	"github.com/alexsuslov/wasty/api/model"
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
	err = model.CreateAdmin(DB)
	if err != nil {
		panic(err)
	}

	//templates
	T, err := Load("templates", ".html")
	if err != nil {
		panic(err)
	}
	validator := auth.Validator(DB)
	basicAuth := auth.BasicAuth(validator)

	r := mux.NewRouter()
	r.HandleFunc("/", hands.Home(DB, T)).Methods("GET")

	r.HandleFunc("/api/v1/drivers",
		basicAuth(hands.Drivers(DB))).Methods("GET")
	r.HandleFunc("/api/v1/cars",
		basicAuth(hands.Cars(DB))).Methods("GET")
	r.HandleFunc("/api/v1/addresses",
		basicAuth(hands.Addresses(DB))).Methods("GET")
	r.HandleFunc("/api/v1/waybills",
		basicAuth(hands.Addresses(DB))).Methods("GET")

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
	filesPath := []string{}
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
