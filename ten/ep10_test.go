package ten

import (
	"testing"
)

func TestSumOfPrimesBelowTwoMillion(t *testing.T) {
	sum := SumOfPrimesBelowTwoMillion()
	if sum != 142913828922 {
		t.Errorf("Expected sum to be 142913828922 but got %v", sum)
	}
}
