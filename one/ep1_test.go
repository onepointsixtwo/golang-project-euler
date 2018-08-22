package one

import (
	"testing"
)

func TestNaturalNumbersInMultiplesOfThreeOrFiveBelowOneThousand(t *testing.T) {
	output := NaturalNumbersInMultiplesOfThreeOrFiveBelowOneThousand()
	if output != 233168 {
		t.Errorf("Unexpected output - expected 233,168 but got %v", output)
	}
}
