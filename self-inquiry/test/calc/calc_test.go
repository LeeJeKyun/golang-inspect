package calc_test

import (
	"go-in-action-self/self-inquiry/test/calc"
	"testing"
)

func TestSum(t *testing.T) {
	s := calc.Sum(1, 2, 3)

	if s != 6 {
		t.Error("Wrong result")
	}
}
