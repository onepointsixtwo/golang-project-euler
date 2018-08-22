package threeeightseven

import (
	"fmt"
	"math"
	"strconv"
)

func sumOfDigitsInNumber(number int) int {
	str := strconv.Itoa(number)
	slice := make([]int, 0)
	for _, char := range str {
		i, err := strconv.Atoi(string(char))
		if err == nil {
			slice = append(slice, i)
		} else {
			fmt.Printf("Error while trying to parse digit %v", err)
		}
	}

	total := 0
	for _, value := range slice {
		total = total + value
	}
	return total
}

func numberIsPrime(potentialPrime int) bool {
	ceilSquareRoot := int(math.Ceil(math.Sqrt(float64(potentialPrime))))

	verifiedNonPrime := false
	for i := 2; i <= ceilSquareRoot; i++ {
		if potentialPrime%i == 0 && i != potentialPrime {
			verifiedNonPrime = true
			break
		}
	}

	return !verifiedNonPrime
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetSumOfStrongRightTruncatableHarshadPrimesLessThanLength(length int) int {
	// All single digit numbers are harshad - their digits add up to themselves, and everything is divisible by at least itself and 1.
	workingHarshadSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	discoveredStrongHarshadNumbers := make([]int, 0)

	//NumDigits starts at two because we already have the first - in the set of starting Harshad numbers
	for numDigits := 1; numDigits < length; numDigits++ {
		updatedHarshadSet := make([]int, 0)

		for _, value := range workingHarshadSet {
			hasStrongPredecessor := contains(discoveredStrongHarshadNumbers, value)

			for digitToAdd := 0; digitToAdd <= 9; digitToAdd++ {
				workingValue := (value * 10) + digitToAdd

				if numberIsPrime(workingValue) && hasStrongPredecessor {
					sum = sum + workingValue
				}

				workingValueSumOfDigits := sumOfDigitsInNumber(workingValue)
				if workingValue%workingValueSumOfDigits == 0 {
					updatedHarshadSet = append(updatedHarshadSet, workingValue)

					potentialStrongPrime := workingValue / workingValueSumOfDigits
					if numberIsPrime(potentialStrongPrime) {
						discoveredStrongHarshadNumbers = append(discoveredStrongHarshadNumbers, workingValue)
					}
				}
			}
		}

		// On the next iteration, we work from the existing harshads'. This is the key really. Since they have to
		// be able to have Harshad parents all the way down, they can only come from that existing set.
		workingHarshadSet = updatedHarshadSet
	}

	return sum
}
