#include <cstdint>
#include <iostream>
#include <string>

using namespace std;

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

int main()
{
    const string s1 = "abc";
    const string s2 = "ahbgdc";

    Solution sol;
    bool k = sol.isSubsequence(s1, s2);

    std::cout << "k = " << k << std::endl;

    return 0;
}