#include <iostream>
#include <string>

using namespace std;

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

int main()
{
    const string s = "   fly me   to   the moon  ";
    Solution sol;
    int k = sol.lengthOfLastWord(s);

    std::cout << "k = " << k << std::endl;

    return 0;
}