# Is Subsequence

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:

```C++
Input: s = "abc", t = "ahbgdc"
Output: true
```

Example 2:

```C++
Input: s = "axc", t = "ahbgdc"
Output: false
```

Constraints:

```C++
0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
```

Follow up:
 Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?

## Solution

### Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

### Approach
<!-- Describe your approach to solving the problem. -->

### Complexity

* Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
Run time <= 0 ms Beats 100.00 % of users with C++

* Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
Memory = 8.64 MB Beat 22.55 % of user with C++

### Code

```C++
class Solution
{
  public:
    bool isSubsequence(string s, string t)
    {
        uint8_t sIndex = 0;
        for (auto it = t.begin(); it != t.end(); ++it)
        {
            if (s[sIndex] == *it)
            {
                ++sIndex;
            }
        }
        return sIndex == s.size();
    }
};
```
