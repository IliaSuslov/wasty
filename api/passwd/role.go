package passwd

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

func (User *Passwd) IsRole(roles ...string) bool {
	for _, r := range User.Roles {
		for _, role := range roles {
			if r == role {
				return true
			}
		}
	}
	return false
}

type Role struct {
	ID   primitive.ObjectID
	Name string
}

func CreateAdmin(DB *mongo.Database) error {
	col := DB.Collection("users")
	res := col.FindOne(context.Background(), bson.M{"name": "admin"})
	if res.Err() == nil {
		return nil
	}
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	password, _ := sid.Generate()
	hash := sha256.Sum256([]byte(password))
	str := fmt.Sprintf("%x", hash)
	hash = sha256.Sum256([]byte(str + "admin"))
	log.Printf("create user admin pass=%v", password)
	_, err := col.InsertOne(context.Background(), &Passwd{
		Name:     "admin",
		Password: hash,
		Roles:    []string{"admin"},
	})
	return err
}


func IsRole(roles interface{}, role string) int {
	Roles := roles.([]string)
	for i, v := range Roles {
		if v == role {
			return i
		}
	}
	return -1
}
