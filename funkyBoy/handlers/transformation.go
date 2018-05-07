package handlers

import (
	"strconv"
	"log"
	"funkyBoy/transformers"
	"fmt"
)

// takes in the map with interface types
// determines which is the underlying type
// of the interface, tries to convert it to float
// and applies MinMaxScaler to it, if the value
// cannot be converted, sets the value to nil
func transform(mp map[string]interface{}) {

	var res float64
	var err error
	for key, val := range  mp {

		// try to convert to float otherwise ignore
		if v, ok := val.(int); ok {

			res = float64(v)
		} else if v, ok := val.(float64); ok {
			res = v
			fmt.Println(res, v)
		} else if  v, ok := val.(string); ok {

			res, err = strconv.ParseFloat(v, 64)
		}
		if err != nil {

			log.Printf("cannot convert [%s] to type float\n", val)
			mp[key] = nil
			// error has to be returned to nill to avoid unexpected behavior
			err = nil
		} else {

			mp[key], err = transformers.Scale(key, res)
		}
	}

}
