package hands

//func Users(DB *mongo.Database, t *template.Template) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		roles := gcontext.Get(r, "user_roles")
//		if model.IsRole(roles, "admin") == -1 {
//			w.WriteHeader(401)
//			return
//		}
//
//		err := t.ExecuteTemplate(w, "users.html", struct{ Title string }{"HOME"})
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}
