# Invert Binary Tree
Given the root of a binary tree, invert the tree, and return its root.

Example 1:

![Example1](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)
```
Input: root = [1,2,2,3,4,4,3]
Output: true
```

Example 2:

![Example2](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)
```
Input: root = [1,2,2,null,3,null,3]
Output: false
```

Example 3:
```
Input: root = []
Output: []
```

Constraints:
```
The number of nodes in the tree is in the range [0, 1000].
-100 <= Node.val <= 100
```

Follow up:
```
Could you solve it both recursively and iteratively?
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->


### Approach
<!-- Describe your approach to solving the problem. -->


### Complexity
- Runtime
 10 ms
Beats 5.52% of users with C++
- Memory
 17.20 MB
Beats 72.03% of users with C++

### Code
```
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        return isSymmetric(root->left, root->right);    
    }
    bool isSymmetric(TreeNode* leftNode, TreeNode* rightNode){
        if (leftNode == nullptr && rightNode == nullptr){
            return true;
        }else if (leftNode == nullptr || rightNode == nullptr){
            return false;
        }else if (leftNode->val != rightNode->val){
            return false;
        }

        return isSymmetric(leftNode->left, rightNode->right) && isSymmetric(leftNode->right, rightNode->left);
    }
};
```