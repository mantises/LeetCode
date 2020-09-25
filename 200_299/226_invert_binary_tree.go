package _200_299

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeByRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = invertTreeByRecursion(root.Right)
	root.Right = invertTreeByRecursion(tmp)
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	return root
}
