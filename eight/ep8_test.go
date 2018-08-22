package eight

import (
	"testing"
)

func TestArrayOfIntegersFromSingleDigitsString(t *testing.T) {
	str := "01234"
	array := ArrayOfIntegersFromSingleDigitsString(str)
	for i := 0; i < 5; i++ {
		if array[i] != i {
			t.Errorf("Value in slice unexpected - expected %v but got %v", i, array[i])
		}
	}
}

func TestGreatestProductAdjacentDigits(t *testing.T) {
	output := GreatestProductAdjacentDigits()
	if output != 23514624000 {
		t.Errorf("Expected output of 23514624000 but got %v", output)
	}
}
