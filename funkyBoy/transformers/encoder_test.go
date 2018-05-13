package transformers

import "testing"

func TestOneHotEncoders(t *testing.T) {
	vl := "A"
	svl := []string{"a", "2", "AJBYB", "0.444"}

	ta := []int{1,0,0,0}

	for i, v := range Encode(svl, vl) {
		if v != ta[i] {
			t.Errorf("Weird stuff: %v", ta[i])
		}
	}
}