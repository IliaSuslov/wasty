package auth

import (
	"github.com/alexsuslov/wasty/api/mid"
	"github.com/alexsuslov/wasty/api/passwd"
	gcontext "github.com/gorilla/context"
	"net/http"
)

type GetUser func(username string, password string) (User *passwd.Passwd, err error)

func BasicAuth(get GetUser) mid.Middle {
	return func(h mid.Handle) mid.Handle {
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
