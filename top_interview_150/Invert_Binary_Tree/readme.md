# Invert Binary Tree
Given the root of a binary tree, invert the tree, and return its root.

Example 1:

![Example1](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)
```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

Example 2:

![Example2](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)
```
Input: root = [2,1,3]
Output: [2,3,1]
```

Example 3:
```
Input: root = []
Output: []
```

Constraints:
```
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
```

Follow up:
```
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->


### Approach
<!-- Describe your approach to solving the problem. -->


### Complexity
- Runtime
0 ms
Beats 100.00% of users with C++
- Memory
11.21 MB
Beats 84.82% of users with C++

### Code
```
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
    void _invertTree(TreeNode* root){
        if (root == nullptr){
            return;
        }
        TreeNode* auxNode = nullptr;
        auxNode = root->left;
        root->left = root->right;
        root->right = auxNode;
        auxNode = nullptr;
        _invertTree(root->left);;
        _invertTree(root->right);
    }

    TreeNode* invertTree(TreeNode* root) {
        _invertTree(root);
        return root;
    }
};
```