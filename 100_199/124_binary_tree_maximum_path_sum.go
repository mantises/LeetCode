package _100_199

import "math"

/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the
sequence has an edge connecting them. A node can only appear in the sequence at most once.
Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any path.

Example 1:
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.


Constraints:
The number of nodes in the tree is in the range [1, 3 * 10^4].
-1000 <= Node.val <= 1000

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(*TreeNode) int
	// 返回以节点 node 为根节点的子树最大路径长度
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		length := node.Val
		if left > 0 {
			length += left
		}
		if right > 0 {
			length += right
		}
		if length > maxSum {
			maxSum = length
		}
		return node.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}
