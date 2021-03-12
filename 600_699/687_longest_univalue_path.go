package _600_699

/*
Given the root of a binary tree, return the length of the longest path, where each node
in the path has the same value. This path may or may not pass through the root.

The length of the path between two nodes is represented by the number of edges between
them.

Example 1:
Input: root = [5,4,5,1,1,5]
Output: 2

Example 2:
Input: root = [1,4,5,4,4,5]
Output: 2


Constraints:
The number of nodes in the tree is in the range [0, 10^4].
-1000 <= Node.val <= 1000
The depth of the tree will not exceed 1000.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	maxLength := 0
	var dfs func(*TreeNode) int
	// 返回以节点 node 为端点延伸出的最长同值路径的长度
	dfs = func(node *TreeNode) int {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		leftLength, rightLength := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			leftLength = left + 1
		}
		if node.Right != nil && node.Val == node.Right.Val {
			rightLength = right + 1
		}
		maxLength = max(maxLength, leftLength+rightLength)
		return max(leftLength, rightLength)
	}
	dfs(root)
	return maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
