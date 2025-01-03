# Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

```C++
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

Example 2:

```C++
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

Constraints:

```C++
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
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
Memory = 11.98 MB Beat 36.86.10 % of user with C++

### Code

```C++
class Solution
{
  protected:
    std::string prefixCommon(const std::string &a, const std::string &b)
    {
        size_t index = 0;
        while (index < a.size() && index < b.size() && a[index] == b[index])
        {
            ++index;
        }
        return a.substr(0, index);
    }

  public:
    string longestCommonPrefix(vector<string> &strs)
    {
        if (strs.empty())
        {
            return "";
        }
        std::string prefix = strs[0];
        for (size_t i = 1; i < strs.size(); ++i)
        {
            prefix = prefixCommon(prefix, strs[i]);
            if (prefix.empty())
            {
                break;
            }
        }
        return prefix;
    }
};
```
