package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"funkyBoy/transformers"
	"strconv"
	"encoding/json"
	"funkyBoy/model"
	"log"
)

// handler function for the scaling operation
// will return the GoodResponse if the scaling was successful
// will return  the BadRequestResponse if the input could not
// be converted to type float64
func GetScaledValue(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fl, err := strconv.ParseFloat(params["value"], 64)

	// if we cannot convert the float then return a 400
	if err != nil {

		log.Println("could not convert input to float, error occured: ", err)

		br := model.NewBadRequestResponse(string(err.Error()))
		jbr, _ := json.Marshal(&br)
		http.Error(w, string(jbr), http.StatusBadRequest)

		return
	}
	scaler := transformers.GetDefaultScaler()
	scaled := scaler.Scale(fl)

	res := model.NewGoodResponse(
		params,
		strconv.FormatFloat(scaled, 'f', 2, 64),
		)
	log.Println("encoding oupt to JSON")
	json.NewEncoder(w).Encode(res)
}