# Convert Sorted Array to Binary Search Tree

Can you solve this real interview question? Convert Sorted Array to Binary Search Tree - Given an integer array `nums` where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

Example 1:

![Example 1](https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg)

```text
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:  
```

![Explanation Image](https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg)

Example 2:

![Example 2](https://assets.leetcode.com/uploads/2021/02/18/btree.jpg)

```text
Input: nums = [1,3]  
Output: [3,1]  
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
```

Constraints:

```text
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums is sorted in a strictly increasing order.
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

### Approach
<!-- Describe your approach to solving the problem. -->

### Complexity

- Runtime
 8 ms
Beats 65.76% of users with C++
- Memory
 22.24 MB
Beats 19.38% of users with C++

### Code

```C++
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode *sortedArrayToBST(std::vector<int> &nums) {
        if (nums.empty()) {
            return nullptr;
        }

        return addNode(nums, 0, nums.size() - 1);
    }

    TreeNode *addNode(std::vector<int> &nums, int start, int end) {

        if (start > end) {
            return nullptr;
        }
        int mid = start + (end - start) / 2;
        TreeNode *node = new TreeNode(nums.at(mid));

        node->left = addNode(nums, start, mid - 1);
        node->right = addNode(nums, mid + 1, end);

        return node;
    }
};
```
