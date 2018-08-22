package seven

import (
	"testing"
)

func TestGetTenThousandAndFirstPrimeNumber(t *testing.T) {
	output := GetTenThousandAndFirstPrimeNumber()
	if output != 104743 {
		t.Errorf("Expected value to be 104743 but was %v", output)
	}
}
