package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	length := 0
	left := 0

	for right, char := range s {
		if lastIndex, found := charIndex[char]; found && lastIndex >= left {
			left = lastIndex + 1
		}
		charIndex[char] = right

		currentLen := right - left + 1

		if currentLen > length {
			length = currentLen
		}
	}

	return length
}

func main() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))
}
