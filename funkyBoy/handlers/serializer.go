package handlers

import (
	"strconv"
	"log"
	"strings"
	"reflect"
)

// 1) convert input type from interface{} to []interface{}
// 2) figure out which type is it (int, float64, string) are currently supported
// 3) convert the value to a type string
// 4) append that string to a slice of type string
// 5) convert the slice of type string to a type string with (,) as a separator
// 6) convert the "slice" string to slice of bytes
func ConvertSliceToBytes(i interface{}) []byte {

	is := reflect.ValueOf(i)
	if is.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// create a slice of interfaces
	s := make([]interface{}, is.Len())
	for i := 0; i < is.Len(); i++ {
		s[i] = is.Index(i).Interface()
	}

	var sl []string
	//find out the type of the input
	for _, vl := range s {
		if r, ok := vl.(int); ok {

			sl = append(sl, strconv.Itoa(r))
		} else if r, ok := vl.(float64); ok {

			sl = append(sl, strconv.FormatFloat(r, 'f', 4, 64))
		} else if r, ok := vl.(string); ok {

			sl = append(sl, r)
		} else {

			log.Fatal("value:", vl, "has an unsupported type and will be avoided")
			panic("unsupported type")
		}
	}
	// convert the slice of strings to a string itself
	sls := strings.Join(sl,",")

	// convert to slice of bytes and return
	return []byte(sls)
}

// 1) convert bytes[] to type string
// 2) convert string to type []string with (,) separators
// 3) figure out the type of each string element
// ("string", "int", "float64") and convert to it
func ConvertBytesToSLice(b []byte) []interface{} {

	// convert to a slice of string
	ss := strings.Split(string(b), ",")

	var r []interface{}
	// populate interface r
	for _, s := range ss {
		if i, err := strconv.Atoi(s); err == nil {
			r = append(r, i)
		} else if f, err := strconv.ParseFloat(s, 64); err == nil {
			r = append(r, f)
		} else {
			r = append(r, s)
		}
	}
	return r
}
