package transformers

import (
	"testing"
	"fmt"
)

func TestBelongsToBin(t *testing.T) {

	st := "[-0.001, 1.0]"

	a := belongsToBin(st, -0.001)

	if a == 0 {
		t.Errorf("should be 1: %v", a)
	}
}

func TestBelongsToBin2(t *testing.T) {

	st := "(-0.001, 1.0]"

	a := belongsToBin(st, -0.001)

	if a == 1 {
		t.Errorf("should be 0: %v", a)
	}
}

func TestBelongsToBin3(t *testing.T) {

	st := "(-0.001, 1.0]"

	a := belongsToBin(st, 1.0)

	if a == 0 {
		t.Errorf("should be 1: %v", a)
	}
}

func TestBelongsToBin4(t *testing.T) {

	st := "(-0.001, 1.0)"

	a := belongsToBin(st, 1.0)

	if a == 1 {
		t.Errorf("should be 0: %v", a)
	}
}


func TestAllocateToBins(t *testing.T) {

	strs := []string{"(-0.001, 1.0]", "(1.0, 2.0]", "(2.0, 3.0]", "(3.0, 4.0]"}

	fmt.Println(AllocateToBins(strs, 2))
}