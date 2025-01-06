#include <cstdint>
#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

class Solution
{
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

int main()
{
    const string s1 = "aa";
    const string s2 = "ab";

    Solution sol;
    bool k = sol.canConstruct(s1, s2);

    std::cout << "k = " << k << std::endl;

    return 0;
}