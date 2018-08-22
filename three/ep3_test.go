package three

import (
	"testing"
)

func TestFindPrimeFactorsForNumber(t *testing.T) {
	largestPrimeFactor := FindLargestPrimeFactorForNumber(600851475143)
	if largestPrimeFactor != 6857 {
		t.Errorf("Expected largest prime factor to be 6857 but was %v", largestPrimeFactor)
	}
}
