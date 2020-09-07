package mid

import "net/http"

type Handle func(w http.ResponseWriter, r *http.Request)
type Middle func(Handle) Handle
type Response func(w http.ResponseWriter)
