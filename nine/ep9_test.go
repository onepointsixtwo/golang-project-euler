package nine

import (
	"testing"
)

func TestPythagoreanTripletsAddingToOneThousandProduct(t *testing.T) {
	output := PythagoreanTripletsAddingToOneThousandProduct()
	if output != 31875000 {
		t.Errorf("Expected 31875000 but got %v", output)
	}
}
