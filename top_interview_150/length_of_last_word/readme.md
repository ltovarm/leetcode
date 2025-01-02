# Lenght of last word

Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring
 consisting of non-space characters only.

Example 1:

```C++
Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
```

Example 2:

```C++
Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
```

Example 3:

```C++
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

Constraints:

```C++
Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.
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
Memory = 8.8 MB Beat 38.77 % of user with C++

### Code

```C++
class Solution
{
  public:
    int lengthOfLastWord(string s)
    {
        int sol = 0;
        for (auto it = s.rbegin(); it != s.rend(); ++it)
        {
            if (*it == ' ' && sol > 0)
            {
                break;
            }
            else if (*it != ' ')
            {
                sol += 1;
            }
        }
        return sol;
    }
};
```
