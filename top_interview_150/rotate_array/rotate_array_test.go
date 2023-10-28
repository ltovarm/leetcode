package rotate_array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		name         string
		nums         []int
		k            int
		expectedNums []int
	}{
		{
			name:         "Test case 1",
			nums:         []int{1, 2, 3, 4, 5, 6, 7},
			k:            3,
			expectedNums: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:         "Test case 2",
			nums:         []int{-1, -100, 3, 99},
			k:            2,
			expectedNums: []int{3, 99, -1, -100},
		},
		{
			name:         "Test case 3",
			nums:         []int{-1},
			k:            2,
			expectedNums: []int{-1},
		},
		{
			name:         "Test case 4",
			nums:         []int{1, 2},
			k:            3,
			expectedNums: []int{2, 1},
		},
		{
			name:         "Test case 5",
			nums:         []int{1, 2, 3, 4, 5, 6, 7},
			k:            9,
			expectedNums: []int{6, 7, 1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rotate(testCase.nums, testCase.k)
			if !reflect.DeepEqual(testCase.nums, testCase.expectedNums) {
				expected := fmt.Sprintf("%v", testCase.expectedNums)
				numsString := fmt.Sprintf("%v", testCase.nums)
				t.Errorf("Expected %s but got %s", numsString, expected)
			}
		})
	}
}
