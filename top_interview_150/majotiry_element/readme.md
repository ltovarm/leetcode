# Majority element

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

Example 1:
```
Input: nums = [3,2,3]
Output: 3
```

Example 2:
```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

Constraints:
```
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
```

Follow-up: Could you solve the problem in linear time and in O(1) space?

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
It was clear to me that I had to focus on a solution that used a map from the beginning. It is true that it is the solution that is going to sacrifice memory.

### Approach
<!-- Describe your approach to solving the problem. -->
Once we have focused on how I was going to solve it, the flow is simple, it updates the number of times a value is repeated. Once finished I have seen a very nice solution in which it takes advantage of a rule that escaped me, that once count passes the half of the slice you can return the value because the solution is unique.

### Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
        Run time <= 16 ms Beats 76.18% of users with Go
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
        Memory = 6.30 MB Beat 26.09% of user with Go
### Code
```
func majorityElement(nums []int) int {
	counts := make(map[int]int)
	result := 0

	for _, value := range nums {
		counts[value]++
		if counts[value] > counts[result] {
			result = value
		}
	}
	return result
}
```