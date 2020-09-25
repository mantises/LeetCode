package _01_99

/*
Given a binary tree, return the inorder traversal of its nodes' values.

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalByTrivial(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	return nil
}
