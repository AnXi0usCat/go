package transformers

import "strings"

// loop over the sl slice to find out
// which value is vl. assign "1" to that
// value and "0: to the res,
// store inside a new slice
func Encode(sl []string, vl string) []int {

	var rs []int
	for _, v := range sl {
		if strings.ToLower(vl) == strings.ToLower(v) {
			rs = append(rs, 1)
		} else {
			rs = append(rs, 0)
		}
	}
	return rs
}


