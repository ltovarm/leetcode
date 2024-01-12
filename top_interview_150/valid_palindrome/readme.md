# Valid palindrome
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```

Example 2:
```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
```

Example 3:
```
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```
Constraints:
```
1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
```

Follow up:
```
```

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
At the beginning I opted for a functional solution, remove everything that is not alpha numeric and work with that string comparing the beginning with the end. Although it worked, I didn't like it in terms of O(n/2) performance. So after some research I developed the second option which is to ignore the elements that are not numerical alpha but without having to preprocess anything.

### Approach
<!-- Describe your approach to solving the problem. -->
 Although it worked, I didn't like it in terms of O(n/2) performance. So after some research I developed the second option which is to ignore the elements that are not numerical alpha but without having to preprocess anything.

### Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
        Run time <= 6 ms Beats  39.38% of users with Go
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
### Code
```
package valid_palindrome

import (
	"strings"
	"unicode"
)

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {

	cad := []rune(s)
	for start, end := 0, len(s)-1; start <= end; {
		if !isAlphanumeric(cad[start]) {
			start++
			continue
		}
		if !isAlphanumeric(cad[end]) {
			end--
			continue
		}
		if strings.ToLower(string(cad[start])) != strings.ToLower(string(cad[end])) {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
```