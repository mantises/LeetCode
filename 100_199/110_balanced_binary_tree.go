package _100_199

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as: a binary tree in which
the left and right subtrees of every node differ in height by no more than 1.


Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Example 3:
Input: root = []
Output: true


Constraints:
The number of nodes in the tree is in the range [0, 5000].
-10^4 <= Node.val <= 10^4

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return height(root) > 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftDepth := height(root.Left)
	rightDepth := height(root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if abs(leftDepth, rightDepth) > 1 {
		return -1
	}
	return 1 + max(leftDepth, rightDepth)
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
