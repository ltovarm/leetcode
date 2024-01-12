package valid_palindrome

import (
	"strings"
	"unicode"
)

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {

	cad := []rune(s)
	for start, end := 0, len(s)-1; start <= end; {
		if !isAlphanumeric(cad[start]) {
			start++
			continue
		}
		if !isAlphanumeric(cad[end]) {
			end--
			continue
		}
		if strings.ToLower(string(cad[start])) != strings.ToLower(string(cad[end])) {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
