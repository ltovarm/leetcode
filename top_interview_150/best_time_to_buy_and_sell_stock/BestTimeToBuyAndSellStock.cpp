#include <iostream>
#include <vector>

using namespace std;

class Solution
{
  public:
    int maxProfit(vector<int> &prices)
    {
        int maxDiff = 0;
        int minSoFar = prices.at(0);
        for (auto value : prices)
        {
            if (value < minSoFar)
            {
                minSoFar = value;
            }
            if (value - minSoFar > maxDiff)
            {
                maxDiff = value - minSoFar;
            }
        }
        return maxDiff;
    }
};

int main()
{
    vector<int> myVector = {7, 6, 4, 3, 1};
    Solution sol;
    int k = sol.maxProfit(myVector);

    std::cout << "k = " << k << std::endl;

    return 0;
}