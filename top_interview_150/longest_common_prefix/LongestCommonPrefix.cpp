#include <iostream>
#include <string>
#include <vector>

using namespace std;

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

int main()
{
    vector<string> s = {"flower", "flow", "flight"};
    Solution sol;
    string k = sol.longestCommonPrefix(s);

    std::cout << "k = " << k << std::endl;

    return 0;
}