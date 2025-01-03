#include <iostream>
#include <string>
#include <vector>

using namespace std;

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

        for (size_t index = 0; index < (haystack.size() - needle.size());
             ++index)
        {
            cout << haystack.substr(index, needle.size()) << endl;
            if (haystack.substr(index, needle.size()) == needle)
            {
                return index;
            }
        }

        return -1;
    }
};

int main()
{
    const string haystack = "abc";
    const string needle = "c";
    Solution sol;
    int k = sol.strStr(haystack, needle);

    std::cout << "k = " << k << std::endl;

    return 0;
}