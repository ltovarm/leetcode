package max_depth_of_binary_tree

import (
	"testing"
)

func TestMaxDepthOfBinaryTree(t *testing.T) {
	testCases := []struct {
		name           string
		root           *TreeNode
		expectedResult int
	}{
		{
			name:           "Test case 1",
			root:           buildTreeFromArray([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			expectedResult: 3,
		},
		{
			name:           "Test case 2",
			root:           buildTreeFromArray([]interface{}{1, nil, 2}),
			expectedResult: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := maxDepth(testCase.root)
			if result != testCase.expectedResult {
				t.Errorf("Expected %d but got %d", testCase.expectedResult, result)
			}
		})
	}
}
