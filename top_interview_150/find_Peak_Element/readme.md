# Find Peak Element

A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.

Example 1:

```
Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
```

Example 2:

```
Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.
```

Constraints:

```
1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i.
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

### Approach
<!-- Describe your approach to solving the problem. -->

### Complexity

- Runtime
 2 ms
Beats 66.23% of users with C++
- Memory
 11.98 MB
Beats 10.59% of users with C++

### Code

```
class Solution {
  public:
    int findPeakElement(vector<int> &nums) {
        if (nums.size() <= 1) {
            return 0;
        } else if (nums.size() == 2){
            if (nums.at(0) > nums.at(1)){
                return 0;
            }
            return 1;
        }
        int sol = binarySearch(nums, 0, nums.size() - 1);
        return sol;
    }

    int binarySearch(vector<int> &nums, int init, int end) {
        int middle;
        while (init <= end) {
            middle = init + (end - init) / 2;
            int prev = 0;
            try {
                prev = nums.at(middle - 1);
            } catch (...) {
                prev = nums.at(middle);
            }
            int post = 0;
            try {
                post = nums.at(middle + 1);
            } catch (...) {
                post = nums.at(middle);
            }

            if ((nums.at(middle) >= prev) && (nums.at(middle) >= post)) {
                return middle;
            } else if (nums.at(middle) < prev) {
                return binarySearch(nums, init, middle - 1);
            } else {
                return binarySearch(nums, middle + 1, end);
            }
        }
        return -1;
    }
};
```
