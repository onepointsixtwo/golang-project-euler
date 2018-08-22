package four

import (
	"testing"
)

func TestStringReverse(t *testing.T) {
	input := "hello"
	expectedOutput := "olleh"
	actualOutput := Reverse(input)
	if actualOutput != expectedOutput {
		t.Errorf("Expected output was %v but actual output was %v", expectedOutput, actualOutput)
	}
}

func TestStringIsPalindrome(t *testing.T) {
	palindromeShort := "I"
	palindromeLong := "tattarrattat"
	nonPalindrome := "abacus"

	shortValue := StringIsPalindrome(palindromeShort)
	longValue := StringIsPalindrome(palindromeLong)
	nonPalindromeValue := StringIsPalindrome(nonPalindrome)

	if !shortValue || !longValue || nonPalindromeValue {
		t.Errorf("Expected short value to be true, was %v, long value to be true, was %v and non palindrome value to be false, was %v", shortValue, longValue, nonPalindromeValue)
	}
}

func TestLargestThreeDigitProductPalindrome(t *testing.T) {
	value := LargestThreeDigitProductPalindrome()
	if value != 906609 {
		t.Errorf("Expected largest three digit product palindrome to be 906609 but was %v", value)
	}
}
