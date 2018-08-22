package seven

import (
	"math"
)

// Copied from three directly, but I kinda want each one to be stand alone in the solution (other than system libs)
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

func GetTenThousandAndFirstPrimeNumber() int {
	generateNextPrime := primesGenerator()
	prime := 0
	for i := 1; i <= 10001; i++ {
		prime = generateNextPrime()
	}
	return prime
}
