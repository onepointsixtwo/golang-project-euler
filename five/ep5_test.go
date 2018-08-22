package five

import (
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	output := CalculateSmallestMultipleOfOneToTwenty()
	if output != 232792560 {
		t.Errorf("Expected output to be 232792560 but was %v", output)
	}
}
