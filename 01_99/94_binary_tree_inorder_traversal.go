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

func inorderTraversalByRecursive(root *TreeNode) []int {
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
	if root == nil {
		return []int{}
	}
	var res []int
	stack := make([]*TreeNode, 0, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, p.Val)
			p = p.Right
		}
	}
	return res
}
