package auth

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/alexsuslov/wasty/api/passwd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validator(DB *mongo.Database) GetUser {
	return func(username string, password string) (User *passwd.Passwd, err error) {
		users := DB.Collection("users")
		ctx := context.Background()
		query := bson.M{
			"name": username,
		}
		User = &passwd.Passwd{}
		err = users.FindOne(ctx, query).Decode(User)
		if err != nil {
			return nil, err
		}
		hash := sha256.Sum256([]byte(password))
		str := fmt.Sprintf("%x", hash)
		hash = sha256.Sum256([]byte(str + username))
		if !bytes.Equal(User.Password[:], hash[:]) {
			return nil, fmt.Errorf("unknown user or wrong password")
		}
		User.Password = [32]byte{}
		return
	}
}
