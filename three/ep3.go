package three

import (
	"math"
)

func numberIsPrime(potentialPrime int) bool {
	ceilSquareRoot := int(math.Ceil(math.Sqrt(float64(potentialPrime))))

	verifiedNonPrime := false
	for i := 2; i <= ceilSquareRoot; i++ {
		// if the potential next prime is divisible by _any_ of this set, it's not a prime
		if potentialPrime%i == 0 && i != potentialPrime {
			verifiedNonPrime = true
			break
		}
	}

	return !verifiedNonPrime
}

func primesGenerator() func() int {

	lastPrime := 1

	return func() int {
		primeFound := false
		potentialNextPrime := lastPrime + 1
		for !primeFound {
			if numberIsPrime(potentialNextPrime) {
				break
			} else {
				potentialNextPrime = potentialNextPrime + 1
			}
		}

		lastPrime = potentialNextPrime
		return potentialNextPrime
	}
}

func FindPrimeFactorsForNumber(number int) []int {
	primeFactors := make([]int, 0)

	generatePrime := primesGenerator()
	workingNumber := number
	currentPrime := generatePrime()
	for !numberIsPrime(workingNumber) {
		if workingNumber%currentPrime == 0 {
			workingNumber = workingNumber / currentPrime
			primeFactors = append(primeFactors, currentPrime)
		} else {
			currentPrime = generatePrime()
		}
	}
	primeFactors = append(primeFactors, workingNumber)

	return primeFactors
}

func FindLargestPrimeFactorForNumber(number int) int {
	largestPrimeFactor := 1
	primeFactors := FindPrimeFactorsForNumber(number)
	for _, factor := range primeFactors {
		if factor > largestPrimeFactor {
			largestPrimeFactor = factor
		}
	}
	return largestPrimeFactor
}
