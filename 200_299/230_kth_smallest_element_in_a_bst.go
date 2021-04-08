package _200_299

/*
Given the root of a binary search tree, and an integer k, return the kth (1-indexed)
smallest element in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1


Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3


Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104


Follow up: If the BST is modified often (i.e., we can do insert and delete operations)
and you need to find the kth smallest frequently, how would you optimize?

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	seq, res := 0, 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		seq += 1
		if seq == k {
			res = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
