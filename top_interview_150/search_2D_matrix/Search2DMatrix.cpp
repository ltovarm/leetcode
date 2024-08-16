#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    int searchTarget(vector<int> &nums, const int target) {
        int init = 0;
        int end = nums.size() - 1;
        int middle;
        while (init <= end) {
            middle = init + (end - init) / 2;

            if (nums.at(middle) == target) {
                return true;
            }

            if (nums.at(middle) < target) {
                init = middle + 1;
            } else {
                end = middle - 1;
            }
        }
        return false;
    }

    bool searchMatrix(vector<vector<int>> &matrix, int target) {
        bool solution = false;
        for (auto array : matrix) {
            solution = searchTarget(array, target);
            if (true == solution) {
                break;
            }
        }
        return solution;
    }
};

int main() {

    std::vector<std::vector<int>> matriz = {
        {1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
    Solution sol;
    bool out = false;
    out = sol.searchMatrix(matriz, 3);
    assert(out);

    std::cout << "target = 3 -> " << out << std::endl;

    out = sol.searchMatrix(matriz, 13);

    assert(!(out));
    std::cout << "target = 13 -> " << out << std::endl;

    matriz = {{1}};

    out = sol.searchMatrix(matriz, 1);
    assert(out);
    std::cout << "target = 1 -> " << out << std::endl;
    return 0;
}