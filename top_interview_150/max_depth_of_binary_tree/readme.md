# Maimum Depth of Binary Tree
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:

![Example1](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)
```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

Example 2:
```
Input: root = [1,null,2]
Output: 2
```

Constraints:
```
The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
```

Follow up:
```
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
The solutions I came up with were to go through the whole tree and increment a couple of indices for each branch or to implement a recursive solution.

### Approach
<!-- Describe your approach to solving the problem. -->
The first option was not only more resource-intensive but also more difficult and ran the risk of losing the focus of what they were asking for. So I opted for the second option.

### Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
        Run time <= 4 ms Beats  69.00% of users with Go
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

### Code
```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)

	return max(depthLeft, depthRight) + 1
}

```