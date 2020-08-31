package model

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct{
	ID primitive.ObjectID
	Name string
	Password [32]byte
	Roles []string
}

func CreateAdmin(DB *mongo.Database)error{
	col := DB.Collection("users")
	res:=col.FindOne(context.Background(), bson.M{"name":"admin"})
	if res.Err()==nil{
		return nil
	}
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	password, _ := sid.Generate()
	hash := sha256.Sum256([]byte(password))
	str := fmt.Sprintf("%x", hash)
	hash = sha256.Sum256([]byte(str+"admin"))
	log.Printf("create user admin pass=%v", password)
	_, err :=col.InsertOne(context.Background(), &User{
		Name: "admin",
		Password: hash,
		Roles: []string{"admin"},
	})
	return err
}

func NewUser(Name string)*User{
	return &User{Name: Name}
}