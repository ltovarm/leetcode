# Merge Sorted Array

You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

 

Example 1:
```
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
```

Example 2:
```
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
```

Example 3:
```
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
```

Constraints:
```
nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109
``` 

Follow up: Can you come up with an algorithm that runs in O(m + n) time?

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
My first idea was to create a slice and complete it while traversing nums1 and nums2 but that is not effective in terms of O(m+n) as the problem calls for. So I looked for a solution with that compromise.

### Approach
<!-- Describe your approach to solving the problem. -->
we go through the slices from m - 1 to 0 and the same for n, as we know that the length of the final slice will be m + n -1 we are going to complete nums1 from the back to the beginning. If necessary, the slice is completed with the remaining nums2.

### Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
        Run time <= 2ms
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
        Memory = 2.3MB

### Code
```
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

    // If there are elements left in nums2, we copy them to nums1.
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
```