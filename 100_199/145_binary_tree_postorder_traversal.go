package _100_199

/*
Given the root of a binary tree, return the post order traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [3,2,1]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,2]
Output: [2,1]


Example 5:
Input: root = [1,null,2]
Output: [2,1]

Constraints:
The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up:
Recursive solution is trivial, could you do it iteratively?
*/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	stack := make([]*TreeNode, 0, 0)
	stack = append(stack, root)
	var pre *TreeNode
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		if (p.Left == nil && p.Right == nil) ||
			(pre != nil && (pre == p.Left || pre == p.Right)) {
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			pre = p
		} else {
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
		}
	}
	return res
}
func postorderTraversalNonRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}
	res = append(res, root.Val)
	return res
}
