package handlers

import (
	"testing"
)

func TestSerializers1(t *testing.T) {

	a := []interface{}{"cats", 0.0, 1000.0, 0, 1}
	c := ConvertBytesToSLice(ConvertSliceToBytes(a))


	for i, _ := range a {
		if a[i] != c[i] {
			t.Errorf("Weird stuff: %v", c[i])
		}
	}
}

func TestSerializers2(t *testing.T) {

	b := []interface{}{ 0, 1000, 0, 1}
	d := ConvertBytesToSLice(ConvertSliceToBytes(b))

	for i, _ := range b {
		if b[i] != d[i] {
			t.Errorf("Weird stuff: %v", d[i])
		}
	}
}