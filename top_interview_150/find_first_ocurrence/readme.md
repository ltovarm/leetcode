# Find the Index of the First Occurrence in a String

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

```C++
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
```

Example 2:

```C++
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
```

Constraints:

```C++
1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
```

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
Memory = 8.86 MB Beat 8.72 % of user with C++

### Code

```C++
class Solution
{
  public:
    int strStr(string haystack, string needle)
    {
        if (haystack == needle)
        {
            return 0;
        }
        else if (haystack.size() < needle.size())
        {
            return -1;
        }

        for (size_t index = 0; index < (haystack.size() - needle.size() + 1);
             ++index)
        {
            if (haystack.substr(index, needle.size()) == needle)
            {
                return index;
            }
        }
        return -1;
    }
};
```
