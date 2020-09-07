package hands

//
//type DriversQuery struct {
//	Name   *string
//	Limit  *int
//	Offset *int
//}
//
//func Drivers(DB *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		user := gcontext.Get(r, "user")
//		User := user.(model.User)
//		if !User.IsRole("admin", "manager") {
//			Denied(w)
//			return
//		}
//		//Query
//		q := r.URL.Query()
//		query := bson.M{}
//		qregex := AddLowerRegex(query, q)
//		qregex("name")
//
//		drivers := DB.Collection("drivers")
//		w.Header().Set("ContentType", "application/json")
//		cur, err := drivers.Find(context.Background(), query, GetOpts(q))
//		if OnError(w, err) {
//			return
//		}
//		res := []*model.Driver{}
//		for cur.Next(context.Background()) {
//			Driver := &model.Driver{}
//			err := cur.Decode(Driver)
//			if OnError(w, err) {
//				return
//			}
//			res = append(res, Driver)
//		}
//		json.NewEncoder(w).Encode(res)
//		if OnError(w, err) {
//			return
//		}
//	}
//}
