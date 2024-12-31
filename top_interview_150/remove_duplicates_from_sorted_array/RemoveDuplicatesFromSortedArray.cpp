#include <iostream>
#include <vector>

using namespace std;

class Solution
{
  public:
    int removeDuplicates(vector<int> &nums)
    {
        vector<int> vectorAux = {nums.at(0)};
        int ret = 1;
        for (auto value : nums)
        {
            if (vectorAux.back() != value)
            {
                vectorAux.push_back(value);
                ret++;
            }
        }
        nums = vectorAux;
        return ret;
    }
};

int main()
{
    vector<int> myVector = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
    Solution sol;
    int k = sol.removeDuplicates(myVector);

    std::cout << "k = " << k << std::endl;
    for (auto value : myVector)
    {
        std::cout << "value = " << value << std::endl;
    }

    return 0;
}