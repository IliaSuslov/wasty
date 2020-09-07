package hands

import "net/http"

type Handle func(w http.ResponseWriter, r *http.Request)
type Middleware func(Handle) Handle
type Response func(w http.ResponseWriter)