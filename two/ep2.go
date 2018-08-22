package two

import (
	"fmt"
)

// Used a generator function - a function which has access to vars it can update and keep track of to generate a sequence.
func createFibonacciGenerator() func() int {
	xMinusOne := 0
	x := 1

	return func() int {
		retVal := x
		nextX := x + xMinusOne
		xMinusOne = x
		x = nextX
		return retVal
	}
}

func EvenFibonacciValuesUnderFourMillion() int {
	generate := createFibonacciGenerator()
	sum := 0
	for {
		nextValue := generate()
		fmt.Printf("Next fibonnacci value is %v\n", nextValue)
		if nextValue <= 4000000 {
			if nextValue%2 == 0 {
				sum += nextValue
			}
		} else {
			break
		}
	}
	return sum
}
