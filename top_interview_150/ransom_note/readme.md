# Ransom Note

Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

Example 1:

```C++
Input: ransomNote = "a", magazine = "b"
Output: false
```

Example 2:

```C++
Input: ransomNote = "aa", magazine = "ab"
Output: false
```

Example 3:

```C++
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

Constraints:

```C++
1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
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
Memory = 11.97 MB Beat 20.58 % of user with C++

### Code

```C++
class Solution {
public:
    bool canConstruct(std::string ransomNote, std::string magazine)
    {
        std::unordered_map<char, int> myHashMap;

        for (char c : magazine)
        {
            myHashMap[c]++;
        }

        for (char c : ransomNote)
        {
            if (myHashMap[c] > 0)
            {
                myHashMap[c]--;
            }
            else
            {
                return false;
            }
        }

        return true;
    }
};
```
