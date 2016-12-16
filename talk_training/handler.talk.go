package talk_training

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/julienschmidt/httprouter"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	response := Talks{
		ID:         1,
		ProductID:  10,
		Message:    "haloo",
		CreateTime: time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	response := "halo"
	json.NewEncoder(w).Encode(response)
}
