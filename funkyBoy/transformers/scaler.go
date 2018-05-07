package transformers

import (
	"log"
	"errors"
)

type Scaler interface  {

	Scale(f float64) float64
}

// scaler type, contains min and max values of the training set
// and min and max values for the scale ranges
type MinMaxScaler struct {

	featureKey string
	minValue float64
	maxValue float64
	minRange float64
	maxRange float64
}

// will scale the input value, based on the
// min and max of the training set, the result
// will range from minRange to maxRange
func (mx MinMaxScaler) Scale(f float64) float64 {

	log.Println("scaling input:", f)
	fStd := (f - mx.minValue) / (mx.maxValue - mx.minValue)
	return fStd * (mx.maxRange - mx.minRange) + mx.minRange
}

// returns a default parametrised MinMaxScaler
// for testing purposes only
func GetDefaultScaler() MinMaxScaler {

	log.Println("Creating a default scaler")
	return MinMaxScaler{
		featureKey: "value",
		minValue:0,
		maxValue:100,
		minRange:0,
		maxRange:1,
	}
}

// goes to the redis data store, and tries to locate values
// based on the supplied key (k), if the search is successful
// it creates a new MinMaxScaler and scales the input float64 (fl)
// it returns scaled (fl) and nil on success
// it returns -9999.99 and Error on failure
func Scale(k string, fl float64) (float64, error) {
	// place holder for redis data store getter implementation
	sl := ds[k]
	if len(sl) != 0 {

		log.Println("creating a new MinMaxScaler with", sl)
		mx :=  MinMaxScaler{k, sl[0],
			sl[1], sl[2], sl[3],
		}
		return mx.Scale(fl), nil
	} else {
		// just in case return som
		return -9999.99, errors.New(k + ": key is missing in Data Store")
	}
}

// our temporary data store
// only use for test purposes
type TempDataStore map[string][]float64

var ds TempDataStore = TempDataStore{
	"first":   []float64{0, 1000, 0, 1},
	"second":  []float64{-0.34, 1000.23, 0, 1},
	"third":   []float64{-1000, 100.3, 0, 1},
	"fourth":  []float64{-1000, 100.3, 0, 1},
	"fifth":   []float64{500, 3463.33, 0, 1},
	"sixth":   []float64{0, 1, 0, 1},
	"seventh": []float64{345, 2333, 0, 1},
	"eight":   []float64{0, 100, 0, 1},
	"ninth":   []float64{-100, 0, 0, 1},
}

func test() {

}



