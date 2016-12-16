package talk_training

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	productID := r.FormValue("product_id")
	productIDInt, errParse := strconv.ParseInt(productID, 10, 64)
	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := GetTalks(productIDInt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	productID := r.FormValue("product_id")
	shopID := r.FormValue("shop_id")
	userID := r.FormValue("user_id")
	message := r.FormValue("message")
	productIDInt, errParse := strconv.ParseInt(productID, 10, 64)
	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shopIDInt, errParse := strconv.ParseInt(shopID, 10, 64)
	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIDInt, errParse := strconv.ParseInt(userID, 10, 64)
	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := InsertTalk(userIDInt, productIDInt, shopIDInt, message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
