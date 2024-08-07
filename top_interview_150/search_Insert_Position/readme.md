# Search Insert Position
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

```
Input: nums = [1,3,5,6], target = 5
Output: 2
```

Example 2:

```
Input: nums = [1,3,5,6], target = 2
Output: 1
```

Example 3:
```
Input: nums = [1,3,5,6], target = 7
Output: 4
```

Constraints:
```
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->


### Approach
<!-- Describe your approach to solving the problem. -->


### Complexity
- Runtime
3 ms 
Beats 71.5% of users with C++
- Memory
12.44 MB
Beats 27.72% of users with C++

### Code
```
class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int init = 0;
        int end = nums.size() - 1;
        int middle;
        while (init <= end) {
            middle = init + (end - init) / 2;

            if (nums.at(middle) == target) {
                return middle;
            }

            if (nums.at(middle) < target) {
                init = middle + 1;
            } else {
                end = middle - 1;
            }
        }
        if (nums.at(middle) <= target) {
            return middle + 1;
        }
        return middle;
    }
};
```