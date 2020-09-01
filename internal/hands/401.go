package hands

import (
	"encoding/json"
	"net/http"
)

func Denied(w http.ResponseWriter) {

	w.WriteHeader(401)
	json.NewEncoder(w).Encode(struct {
		Success bool
		Message string
	}{
		false,
		"access denied",
	})
}