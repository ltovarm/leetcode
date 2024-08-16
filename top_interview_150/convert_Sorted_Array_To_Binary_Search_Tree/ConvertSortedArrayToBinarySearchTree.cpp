#include <iostream>
#include <queue>
#include <vector>

// Defstartion for a binary tree node.
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right)
        : val(x), left(left), right(right) {}
};

class Solution {
  public:
    TreeNode *sortedArrayToBST(std::vector<int> &nums) {
        if (nums.empty()) {
            return nullptr;
        }

        return addNode(nums, 0, nums.size() - 1);
    }

    TreeNode *addNode(std::vector<int> &nums, int start, int end) {

        if (start > end) {
            return nullptr;
        }
        int mid = start + (end - start) / 2;
        TreeNode *node = new TreeNode(nums.at(mid));

        node->left = addNode(nums, start, mid - 1);
        node->right = addNode(nums, mid + 1, end);

        return node;
    }

    void printBST(TreeNode *root) {
        if (nullptr == root) {
            std::cout << " ,-1 " << std::endl;
            return;
        }

        std::queue<TreeNode *> q;
        q.push(root);

        while (!q.empty()) {
            TreeNode *node = q.front();
            q.pop();

            if (nullptr == node) {
                std::cout << " ,-1 " << std::endl;
            } else {
                std::cout << " ," << node->val << std::endl;
                q.push(node->left);
                q.push(node->right);
            }
        }
    }
};

int main() {
    std::vector<int> nums = {-10, -3, 0, 5, 9};

    Solution sol;

    sol.printBST(sol.sortedArrayToBST(nums));

    return 0;
};