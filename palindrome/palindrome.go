package palindrome

import (
	"strings"
	"unicode"
)

func Palindrome(word string) bool {
	left := 0
	right := len(word) - 1
	lowerWord := strings.ToLower(word)
	for left < right {
		if !unicode.IsDigit(rune(lowerWord[left])) && !unicode.IsLetter(rune(lowerWord[left])) {
			left++
			continue

		}

		if !unicode.IsDigit(rune(lowerWord[right])) && !unicode.IsLetter(rune(lowerWord[right])) {
			right--
			continue

		}
		if lowerWord[left] != lowerWord[right] {
			return false

		}
		left++
		right--
	}

	return true
}

//0(n)
