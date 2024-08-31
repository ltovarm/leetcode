#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    int findPeakElement(vector<int> &nums) {
        if (nums.size() <= 1) {
            return 0;
        }
        int sol = binarySearch(nums, 0, nums.size() - 1);
        return sol;
    }

    int binarySearch(vector<int> &nums, int init, int end) {
        int middle;
        while (init <= end) {
            middle = init + (end - init) / 2;
            int prev = 0;
            try {
                prev = nums.at(middle - 1);
            } catch (...) {
                prev = nums.at(middle);
            }
            int post = 0;
            try {
                post = nums.at(middle + 1);
            } catch (...) {
                post = nums.at(middle);
            }

            if ((nums.at(middle) >= prev) && (nums.at(middle) >= post)) {
                return middle;
            } else if (nums.at(middle) < prev) {
                return binarySearch(nums, init, middle - 1);
            } else {
                return binarySearch(nums, middle + 1, end);
            }
        }
        return -1;
    }
};

void comprobar(vector<int> &nums) {

    Solution sol;
    int resultado = sol.findPeakElement(nums);

    if (resultado != -1)
        std::cout << "Elemento encontrado en el índice " << resultado
                  << std::endl;
    else
        std::cout << "Elemento no encontrado en el arreglo." << std::endl;
};

int main() {

    std::vector<int> nums = {1, 2, 3, 1};

    comprobar(nums);

    nums = {1, 2, 1, 3, 5, 6, 4};

    comprobar(nums);

    nums = {2, 1};

    comprobar(nums);

    return 0;
}