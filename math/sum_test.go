package math

import (
	"testing"
)

func TestSumm(t *testing.T) {

	sum := Sum(4, 5)
	expected := 9

	if sum != expected {
		t.Errorf("got %q, wanted %q", sum, expected)
	}
}
