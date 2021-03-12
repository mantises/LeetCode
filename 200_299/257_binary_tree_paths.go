package _200_299

import "strconv"

/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:
Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]
Explanation: All root-to-leaf paths are: 1->2->5, 1->3

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	var res []string
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, prefix string) {
		if prefix != "" {
			prefix += "->"
		}
		s := prefix + strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, s)
		} else {
			if node.Left != nil {
				dfs(node.Left, s)
			}
			if node.Right != nil {
				dfs(node.Right, s)
			}
		}
	}
	dfs(root, "")
	return res
}
