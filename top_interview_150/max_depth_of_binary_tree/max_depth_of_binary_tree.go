package max_depth_of_binary_tree

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeFromArray(array []interface{}) *TreeNode {
	return buildTree(array, 0)
}

func buildTree(array []interface{}, index int) *TreeNode {
	if index >= len(array) || array[index] == nil {
		return nil
	}

	value, ok := array[index].(int)
	if !ok {
		// Handle the case where the value is not an integer
		fmt.Println("Error: Value is not an integer.")
		return nil
	}

	node := &TreeNode{Val: value}
	node.Left = buildTree(array, 2*index+1)
	node.Right = buildTree(array, 2*index+2)

	return node
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)

	return max(depthLeft, depthRight) + 1
}
