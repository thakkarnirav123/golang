package utility

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	testPalindromeVars := []struct {
		word   string
		result bool
	}{
		{"", false},
		{"abc", false},
		{"abcdcba", true},
		{"abcd", false},
		{"abcddcba", true},
		{"abcddcba ", false},
		{" abcddcba", false},
		{" abcddcba ", true},

	}
	for _, testPalindromeVar := range testPalindromeVars {
		t.Run(testPalindromeVar.word, func(t *testing.T) {

			result := isPalindrome(testPalindromeVar.word)

			if result != testPalindromeVar.result {
				t.Errorf("%s : want %v result %v", testPalindromeVar.word, result, testPalindromeVar.result)
			} else {
				fmt.Print(testPalindromeVar.word, " is palindrome")
			}
		})
	}
}
