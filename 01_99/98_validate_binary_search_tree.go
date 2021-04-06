package _01_99

import "math"

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


Constraints:
The number of nodes in the tree is in the range [1, 10^4].
-2^31 <= Node.val <= 2^31 - 1

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return validHelper(root, math.MinInt64, math.MaxInt64)
}

func validHelper(root *TreeNode, low, high int) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if root.Val <= low || root.Val >= high {
		return false
	}
	return validHelper(root.Left, low, root.Val) && validHelper(root.Right, root.Val, high)
}

func isValidBSTByInOrderTraversal(root *TreeNode) bool {
	pre := math.MinInt64
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}
