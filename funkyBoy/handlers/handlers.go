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
func GetScaledValueTest(w http.ResponseWriter, r *http.Request) {

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
	log.Println("encoding output to JSON")
	json.NewEncoder(w).Encode(res)
}


func GetScaledValue(w http.ResponseWriter, r *http.Request){

	decoder := json.NewDecoder(r.Body)
	// make it of type interface so it could accept any value
	req := make(map [string]interface{})

	decoder.Decode(&req)
	transform(req)

	resp := make(map [string]interface{})
	resp["code"] = 200
	resp["data"] = req
	json.NewEncoder(w).Encode(resp)
}


