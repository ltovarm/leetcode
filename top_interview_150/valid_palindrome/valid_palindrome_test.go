package valid_palindrome

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name           string
		s              string
		expectedResult bool
	}{
		{
			name:           "Test case 1",
			s:              "A man, a plan, a canal: Panama",
			expectedResult: true,
		},
		{
			name:           "Test case 2",
			s:              "race a car",
			expectedResult: false,
		},
		{
			name:           "Test case 3",
			s:              " ",
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isPalindrome(testCase.s)
			if result != testCase.expectedResult {
				t.Errorf("Expected %t but got %t", testCase.expectedResult, result)
			}
		})
	}
}
