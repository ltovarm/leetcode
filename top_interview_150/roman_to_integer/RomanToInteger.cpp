#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

class Solution
{
  public:
    int romanToInt(string s)
    {
        int total = 0;
        int prevValue = 0;

        for (auto it = s.rbegin(); it != s.rend(); ++it)
        {
            int value = romanToIntMap[*it];
            if (value < prevValue)
            {
                total -= value;
            }
            else
            {
                total += value;
            }
            prevValue = value;
        }

        return total;
    }

  private:
    unordered_map<char, int> romanToIntMap = {
        {'I', 1},   {'V', 5},   {'X', 10},  {'L', 50},
        {'C', 100}, {'D', 500}, {'M', 1000}};
};

int main()
{
    const string s = "MCMXCIV";
    Solution sol;
    int k = sol.romanToInt(s);

    std::cout << "k = " << k << std::endl;

    return 0;
}