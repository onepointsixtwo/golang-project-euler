package four

import (
	"strconv"
)

func Reverse(str string) string {
	if str != "" {
		return Reverse(str[1:]) + str[:1]
	}
	return ""
}

func StringIsPalindrome(value string) bool {
	return value == Reverse(value)
}

func LargestThreeDigitProductPalindrome() int {
	lowValue := 100
	highValue := 999
	palindromes := make([]int, 0)
	for i := highValue; i >= lowValue; i-- {
		for x := highValue; x >= lowValue; x-- {
			currentValue := i * x
			stringRepresentation := strconv.Itoa(currentValue)
			if StringIsPalindrome(stringRepresentation) {
				palindromes = append(palindromes, currentValue)
			}
		}
	}

	max := 0
	for _, palindrome := range palindromes {
		if palindrome > max {
			max = palindrome
		}
	}
	return max
}
