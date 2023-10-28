# Remove Element

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

 

Example 1:
```
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

Example 2:
```
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
``` 

Constraints:
```
0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
Actually, a little trickery has been done as it is not necessary to remove the elements, only to add the undeleted elements to the beginning of the slice.

### Approach
<!-- Describe your approach to solving the problem. -->
The solution is very simple, it goes through the array and if the element is different to the value that you want to "eliminate" it is reassigned in the slice.

### Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
        Run time <= 2ms Beats 65.72 % of users with Go
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
        Memory = 2.19 MB Beat 13.17% of user with Go
### Code
```
func removeElement(nums []int, val int) int {

	// Initialize values
	result := 0

	// Goes through the array and counts the elements that match val
	// and modifies the nums slice.
	for _, value := range nums {
		if value != val {
			result++
			nums[result] = value
		}
	}

	return result
}
```