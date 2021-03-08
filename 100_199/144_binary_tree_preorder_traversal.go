package _100_199

/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.
Example 1:
Input: root = [1,null,2,3]
Output: [1,2,3]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,2]
Output: [1,2]

Example 5:
Input: root = [1,null,2]
Output: [1,2]

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up:
Recursive solution is trivial, could you do it iteratively?
*/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

func preorderTraversalNonRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	stack := make([]*TreeNode, 0, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			res = append(res, p.Val)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return res
}
