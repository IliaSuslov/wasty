package hands

type CarsQuery struct {
	Name   *string
	Limit  *int
	Offset *int
}

//func Cars(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//rights
//		user := gcontext.Get(r, "user")
//		User := user.(model.User)
//		if !User.IsRole("admin", "manager") {
//			Denied(w)
//			return
//		}
//
//		//query
//		q := r.URL.Query()
//		query := bson.M{}
//		qregex := AddLowerRegex(query, q)
//		qregex("name")
//
//		c := DB.Collection("cars")
//		w.Header().Set("ContentType", "application/json")
//		//todo: limit, skip
//		cur, err := c.Find(context.Background(), query, GetOpts(q))
//		if OnError(w, err) {
//			return
//		}
//		res := []*model.Car{}
//		for cur.Next(context.Background()) {
//			Car := &model.Car{}
//			err := cur.Decode(Car)
//			if OnError(w, err) {
//				return
//			}
//			res = append(res, Car)
//		}
//		json.NewEncoder(w).Encode(res)
//		if OnError(w, err) {
//			return
//		}
//	}
//}

//func Car(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//rights
//		user := gcontext.Get(r, "user")
//		User := user.(model.User)
//		if !User.IsRole("admin", "manager") {
//			Denied(w)
//			return
//		}
//
//		//query
//		q := r.URL.Query()
//		query := bson.M{}
//		qregex := AddLowerRegex(query, q)
//		qregex("name")
//
//		c := DB.Collection("cars")
//		w.Header().Set("ContentType", "application/json")
//		//todo: limit, skip
//		cur, err := c.Find(context.Background(), query, GetOpts(q))
//		if OnError(w, err) {
//			return
//		}
//		res := []*model.Car{}
//		for cur.Next(context.Background()) {
//			Car := &model.Car{}
//			err := cur.Decode(Car)
//			if OnError(w, err) {
//				return
//			}
//			res = append(res, Car)
//		}
//		json.NewEncoder(w).Encode(res)
//		if OnError(w, err) {
//			return
//		}
//	}
//}
