package nine

import (
	"fmt"
	"math"
)

func PythagoreanTripletsAddingToOneThousandProduct() int {
	for a := 2; a < 1000; a++ {
		for b := 998; b > 0; b-- {
			c := 1000 - a - b
			// Do abc satisfy the equation? If so return a*b*c
			// b is hypoteneuse - answer should be found before the loops make it anything other than the largest value of the three
			hypSquared := int(math.Pow(float64(b), 2))
			othersSquared := int(math.Pow(float64(a), 2) + math.Pow(float64(c), 2))
			if hypSquared == othersSquared {
				fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
				return a * b * c
			}
		}
	}
	return 0
}
