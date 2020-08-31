package auth

import (
	"github.com/gorilla/context"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)
type GetRoles func(username string, password string) (roles []string, err error)

func BasicAuth(get GetRoles) func(handler) handler {
	return func(h handler) handler {
		return func(w http.ResponseWriter, r *http.Request) {
			user, pass, _ := r.BasicAuth()
			if user == "" {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"MY REALM\"")
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
				return
			}
			roles, err := get(user, pass)
			if err != nil {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"MY REALM\"")
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
				return
			} else {
				context.Set(r, "user_roles", roles)
				//w.Header().Set("w_roles", strings.Join(roles, ","))
			}
			h(w, r)
		}
	}

}
