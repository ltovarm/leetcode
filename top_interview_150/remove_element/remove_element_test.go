package remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		name           string
		nums           []int
		val            int
		expectedResult int
		expectedNums   []int
	}{
		{
			name:           "Test case 1",
			nums:           []int{3, 2, 2, 3},
			val:            3,
			expectedResult: 2,
			expectedNums:   []int{2, 2},
		},
		{
			name:           "Test case 2",
			nums:           []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:            2,
			expectedResult: 5,
			expectedNums:   []int{0, 1, 3, 0, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := removeElement(testCase.nums, testCase.val)
			if result != testCase.expectedResult {
				t.Errorf("Expected %d but got %d", testCase.expectedResult, result)
			}
		})
	}
}
