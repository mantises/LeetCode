package _400_499

/*
Find the sum of all left leaves in a given binary tree.

Example:

3
/  \
9   20
  /   \
  15    7

There are two left leaves in the binary tree, with values 9 and 15 respectively
return 24.
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum = root.Left.Val
	}
	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

// TODO solve this problem by DFS & WFS
func sumOfLeftLeavesDFS(root *TreeNode) int {
	sum := 0
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil && r.Left.Left == nil && r.Left.Right == nil {
			sum += r.Left.Val
		}
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return sum
}
