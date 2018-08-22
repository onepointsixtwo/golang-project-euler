package threeeightseven

import (
	"testing"
)

func TestSumOfDigitsInNumber(t *testing.T) {
	if sumOfDigitsInNumber(100) != 1 {
		t.Error("The sum of digits in 100 should be 1")
	}
	if sumOfDigitsInNumber(2833) != 16 {
		t.Error("The sum of digits in 2833 should be 16")
	}
	if sumOfDigitsInNumber(180000000) != 9 {
		t.Error("The sub of digits in 180000000 should be 9")
	}
}

func TestGetSumOfStrongRightTruncatableHarshadPrimesUnderTenPowFourteen(t *testing.T) {
	output := GetSumOfStrongRightTruncatableHarshadPrimesLessThanLength(14)
	if output != 47175350800711 {
		t.Errorf("Expected output was 47175350800711 but got %v", output)
	}
}
