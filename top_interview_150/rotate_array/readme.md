# Rotate Array

Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

 

Example 1:
```
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

Example 2:
```
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
```

Constraints:
```
1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
```

Follow up:
```
Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
At first I thought of a brute force solution as usual. Once implemented and working I decided to improve the code looking for a better solution in time and memory.

### Approach
<!-- Describe your approach to solving the problem. -->
The solution is to reorder the whole array and then divide it in 2, the part that would be at the beginning and the final part once you know these two divisions, depending on k, both parts are inverted again. 
Note that to rotate it to the left instead of to the right you have to rotate first the divided slices and finally the whole slice.

PS: you have put a sentence to avoid unnecessary rotations when k > length. For example, if you have to rotate an array of 7 elements 9 times, it is the same as rotating it 2 times.

### Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
        Run time <= 29 ms Beats  71.31% of users with Go
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
        Memory = 7.23 MB Beat 97.41% of user with Go
### Code
```
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {

	lenght := len(nums)

	// Adjust n in case k > lenght avoid incessant rotations
	k = k % lenght

	reverse(nums, 0, lenght-1) // Reverse entire array
	reverse(nums, 0, k-1)      // Reverse first half of array
	reverse(nums, k, lenght-1) // Reverse second half of array
}
```