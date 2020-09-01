package auth

import (
	"github.com/alexsuslov/wasty/api/model"
	gcontext "github.com/gorilla/context"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)
type GetUser func(username string, password string) (user *model.User, err error)

func BasicAuth(get GetUser) func(handler) handler {
	return func(h handler) handler {
		return func(w http.ResponseWriter, r *http.Request) {
			user, pass, _ := r.BasicAuth()
			if user == "" {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"MY REALM\"")
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
				return
			}
			User, err := get(user, pass)
			if err != nil {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"MY REALM\"")
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
				return
			} else {
				gcontext.Set(r, "user", User)
			}
			h(w, r)
		}
	}

}
