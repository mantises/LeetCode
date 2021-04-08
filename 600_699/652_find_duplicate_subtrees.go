package _600_699

import "strconv"

/*
Given the root of a binary tree, return all duplicate subtrees.

For each kind of duplicate subtrees, you only need to return the root node of any one
of them.

Two trees are duplicate if they have the same structure with the same node values.

Example 1:
Input: root = [1,2,3,4,null,2,4,null,null,4]
Output: [[2,4],[4]]

Example 2:
Input: root = [2,1,1]
Output: [[1]]

Example 3:
Input: root = [2,2,2,3,null,3,null]
Output: [[2,3],[3]]


Constraints:
The number of the nodes in the tree will be in the range [1, 10^4]
-200 <= Node.val <= 200

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	hash := make(map[string]int, 0)
	res := make([]*TreeNode, 0)
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		str := strconv.Itoa(node.Val) + "," + dfs(node.Left) + "," + dfs(node.Right)
		if hash[str] == 1 {
			res = append(res, node)
		}
		hash[str] += 1
		return str
	}
	dfs(root)
	return res
}
