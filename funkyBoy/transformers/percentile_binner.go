package transformers

import (
	"strings"
	"strconv"
)

// loop over the slice of bin intervals and
// retrun the binned slice
func AllocateToBins(sl []string, f float64) []int  {

	var res []int
	for _, vl := range sl {
		res  = append(res, belongsToBin(vl, f))
	}
	return res
}

//
// 1) check the first and the last element of
// the string to figure out if the first/last
// elements are inclusive or exclusive (-0.001, 1.0]
//
// 2) extract the minimum and the maxim
// value for the particular bin
//
// 3) check if the input value belongs to the current bin
func belongsToBin(s string, f float64) int {

	// get the first and last elements
	ft := string(s[0])
	lt := string(s[len(s) - 1])

	// remove first and last elements from the string
	s = s[1 : len(s) - 1]

	// remove white spaces from the string
	s = strings.Replace(s, " ", "", -1)

	// split the string based onn the comma
	ss := strings.Split(s, ",")

	//convert to floating point numbers
	fl1, err := strconv.ParseFloat(ss[0], 64)

	if err != nil {
		panic("Cannot convert the first element fo float64")
	}
	 fl2, err := strconv.ParseFloat(ss[1], 64)

	if err != nil {
		panic("Cannot convert the second element fo float64")
	}

	var fc bool
	var sc bool

	// check the first condition
	switch ft {

	case "[":
		fc = f >= fl1
	case "(":
		fc = f > fl1
	default:
		fc = false
	}

	// check second condition
	switch lt {

	case "]":
		sc = f <= fl2
	case ")":
		sc = f < fl2
	default:
		sc = false
	}

	// if the input is in between then return 1
	if fc == true && sc == true {
		return 1
	}
	return 0
}
