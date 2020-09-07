package hands

import (
	"encoding/json"
	"log"
	"net/http"
)

func OnError(w http.ResponseWriter, err error) bool {

	if err == nil {
		return false
	}
	log.Println(err)
	w.WriteHeader(500)
	err = json.NewEncoder(w).Encode(struct {
		Success bool
		Message string
	}{
		false,
		err.Error(),
	})
	if err != nil {
		log.Println(err)
	}
	return true
}
