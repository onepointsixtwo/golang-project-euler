package six

import (
	"math"
)

func GetSumSquareDifferenceForFirstOneHundredNaturalNumbers() int {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
		sumOfSquares = sumOfSquares + int(math.Pow(float64(i), 2))
	}

	squareOfSum := int(math.Pow(float64(sum), 2))
	return squareOfSum - sumOfSquares
}
