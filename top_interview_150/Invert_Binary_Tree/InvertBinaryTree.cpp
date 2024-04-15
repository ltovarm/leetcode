#include <iostream>
using namespace std;

//  Definition for a binary tree node.
struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode() : val(0), left(nullptr), right(nullptr) {}
     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 };
 
class Solution {
public:
    void _invertTree(TreeNode* root){
        if (root == nullptr){
            return;
        }
        TreeNode* auxNode = nullptr;
        auxNode = root->left;
        root->left = root->right;
        root->right = auxNode;
        auxNode = nullptr;
        _invertTree(root->left);;
        _invertTree(root->right);
    }

    TreeNode* invertTree(TreeNode* root) {
        _invertTree(root);
        return root;
    }
};



int main() {

    return 0;
}