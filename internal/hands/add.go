package hands

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/url"
	"strconv"
	"strings"
)

//AddStr url value to bson query
func AddStr(query bson.M, q url.Values) func(string) {
	return func(name string) {
		if q.Get(name) != "" {
			query[strings.ToLower(name)] = q.Get(name)
		}
	}
}

//AddInt url value to bson query
func AddInt(query bson.M, q url.Values) func(string) {
	return func(name string) {
		if q.Get(name) != "" {
			v, err := strconv.Atoi(q.Get(name))
			if err != nil {
				v = 0
			}
			query[strings.ToLower(name)] = v
		}
	}
}

//AdderLower url value to bson query
func AddLower(query bson.M, q url.Values) func(string) {
	return func(name string) {
		if q.Get(name) != "" {
			query[strings.ToLower(name)] = q.Get(name)
		}
	}
}

//AdderLowerRegex url value to bson query
func AddLowerRegex(query bson.M, q url.Values) func(string) {
	return func(name string) {
		if q.Get(name) != "" {
			strings.ToLower(name)
			query[strings.ToLower(name)] = bson.M{"$regex": q.Get(name)}
		}
	}
}
