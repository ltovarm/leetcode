#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    int searchInsert(vector<int> &nums, int target) {
        int init = 0;
        int end = nums.size() - 1;
        int middle;
        while (init <= end) {
            middle = init + (end - init) / 2;

            if (nums.at(middle) == target) {
                return middle;
            }

            if (nums.at(middle) < target) {
                init = middle + 1;
            } else {
                end = middle - 1;
            }
        }
        if (nums.at(middle) <= target) {
            return middle + 1;
        }
        return middle;
    }
};

void comprobar(vector<int> &nums, int target) {

    Solution sol;
    int resultado = sol.searchInsert(nums, target);

    if (resultado != -1)
        std::cout << "Elemento encontrado en el índice " << resultado
                  << std::endl;
    else
        std::cout << "Elemento no encontrado en el arreglo." << std::endl;
};

int main() {

    std::vector<int> nums = {1, 3, 5, 6};
    int target = 5;

    comprobar(nums, target);

    nums = {1, 3, 5, 6};
    target = 2;

    comprobar(nums, target);

    nums = {1, 3, 5, 6};
    target = 7;

    comprobar(nums, target);
    return 0;
}