package majority

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		name           string
		nums           []int
		expectedResult int
	}{
		{
			name:           "Test case 1",
			nums:           []int{3, 2, 3},
			expectedResult: 3,
		},
		{
			name:           "Test case 2",
			nums:           []int{2, 2, 1, 1, 1, 2, 2},
			expectedResult: 2,
		},
		{
			name:           "Test case 3",
			nums:           []int{6, 5, 5},
			expectedResult: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := majorityElement(testCase.nums)
			if result != testCase.expectedResult {
				t.Errorf("Expected %d but got %d", testCase.expectedResult, result)
			}
		})
	}
}
