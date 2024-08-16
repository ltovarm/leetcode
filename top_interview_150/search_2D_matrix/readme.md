# Search a 2D Matrix
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
* The first integer of each row is greater than the last integer of the previous row.
* Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:

![Example1](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)
```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

Example 2:

![Example2](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)
```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

Constraints:
```
* m == matrix.length
* n == matrix[i].length
* 1 <= m, n <= 100
* -104 <= matrix[i][j], target <= 104
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
As I had already worked with a binary search before, I knew that this was the solution. I just had to do an interrogator and a binary search on each array.

### Approach
<!-- Describe your approach to solving the problem. -->
Reusing code

### Complexity
- Runtime
0 ms
Beats 100% of users with C++
- Memory
12.38 MB
Beats 25.31% of users with C++

### Code
```
class Solution {
  public:
    int searchTarget(vector<int> &nums, const int target) {
        int init = 0;
        int end = nums.size() - 1;
        int middle;
        while (init <= end) {
            middle = init + (end - init) / 2;

            if (nums.at(middle) == target) {
                return true;
            }

            if (nums.at(middle) < target) {
                init = middle + 1;
            } else {
                end = middle - 1;
            }
        }
        return false;
    }

    bool searchMatrix(vector<vector<int>> &matrix, int target) {
        bool solution = false;
        for (auto array : matrix) {
            solution = searchTarget(array, target);
            if (true == solution) {
                break;
            }
        }
        return solution;
    }
};
```