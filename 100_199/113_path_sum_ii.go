package _100_199

/*
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths
where each path's sum equals targetSum.

A leaf is a node with no children.


Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: []

Example 3:
Input: root = [1,2], targetSum = 0
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000


*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var dfs func(node *TreeNode, prePath []int, target int)
	dfs = func(node *TreeNode, prePath []int, target int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && node.Val == target {
			var dst = make([]int, len(prePath))
			copy(dst[:], prePath)
			res = append(res, append(dst, target))
		} else {
			prePath = append(prePath, node.Val)
			dfs(node.Left, prePath, target-node.Val)
			dfs(node.Right, prePath, target-node.Val)
		}
	}
	dfs(root, []int{}, targetSum)
	return res
}
