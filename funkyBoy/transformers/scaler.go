package transformers

import "log"

type Scaler interface  {

	Scale(f float64) float64
}

// scaler type, contains min and max values of the training set
// and min and max values for the scale ranges
type MinMaxScaler struct {

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
func GetDefaultScaler() MinMaxScaler {

	log.Println("Creating a default scaler")
	return MinMaxScaler{
		minValue:0,
		maxValue:100,
		minRange:0,
		maxRange:1,
	}
}
