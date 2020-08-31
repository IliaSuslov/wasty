package auth

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/alexsuslov/wasty/api/model"
	"crypto/sha256"
)

func Validator(DB *mongo.Database)GetRoles{
	return func(username string, password string)(roles []string, err error){
		users := DB.Collection("users")
		ctx:= context.Background()
		query:=bson.M{
			"name":username,
		}
		User:=&model.User{}
		err = users.FindOne(ctx, query).Decode(User)
		if err!=nil{
			return nil, err
		}
		hash := sha256.Sum256([]byte(password))
		str := fmt.Sprintf("%x", hash)
		hash = sha256.Sum256([]byte(str+username))
		if !bytes.Equal(User.Password[:], hash[:]){
			return nil, fmt.Errorf("unknown user or wrong password")
		}
		return User.Roles, nil
	}
}