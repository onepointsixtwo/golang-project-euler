package six

import (
	"testing"
)

func TestGetSumSquareDifferenceForFirstOneHundredNaturalNumbers(t *testing.T) {
	output := GetSumSquareDifferenceForFirstOneHundredNaturalNumbers()
	if output != 25164150 {
		t.Errorf("Expected output to be 25164150 but was %v", output)
	}
}
