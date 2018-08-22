package two

import (
	"testing"
)

func TestEvenFibonacciValuesUnderFourMillion(t *testing.T) {
	value := EvenFibonacciValuesUnderFourMillion()
	if value != 4613732 {
		t.Errorf("Expected output to be 4,613,732 but it was %v", value)
	}
}
